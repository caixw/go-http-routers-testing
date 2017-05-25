// Copyright 2017 by caixw, All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package data

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"runtime"
	"strconv"
	"testing"

	"github.com/caixw/go-http-routes-testing/apis"
	"github.com/caixw/go-http-routes-testing/routers"
)

// JSON 输出 JSON 数据
func JSON(dir string) {
	for index, c := range apis.APIS {
		items := make([]*item, 0, len(routers.Routers))

		for _, r := range routers.Routers {
			items = append(items, single(c, r))
		}

		bs, err := json.MarshalIndent(items, "", "  ")
		if err != nil {
			panic(err)
		}

		path := filepath.Join(dir, strconv.Itoa(index)+".json")
		f, err := os.Create(path)
		if err != nil {
			panic(err)
		}
		defer f.Close()
		f.Write(bs)
	}
}

// HTML 输出 HTML 模板
func HTML(dir string) {
	//
}

func single(c *apis.Collection, r *routers.Router) *item {
	ret := &item{
		RouterName: r.Name,
		APIName:    c.Name,
	}

	// 加载
	h, size := loadAPIS(c.APIS, r.Load)
	ret.MemBytes = size

	// 命中情况
	ret.Hits = testHit(c.APIS, h)

	// bench
	rslt := testing.Benchmark(func(b *testing.B) {
		b.ReportAllocs()
		b.ResetTimer()

		for i := 0; i < b.N; i++ {
			api := c.APIS[i%len(c.APIS)]

			w := httptest.NewRecorder()
			r := httptest.NewRequest(api.Method, api.Test, nil)
			h.ServeHTTP(w, r)
		}
	})
	ret.Bench = &bench{
		NsPerOp:           rslt.NsPerOp(),
		AllocsPerOp:       rslt.AllocsPerOp(),
		AllocedBytesPerOp: rslt.AllocedBytesPerOp(),
	}

	return ret
}

func testHit(apis []*apis.API, h http.Handler) []*hit {
	w := httptest.NewRecorder()
	hits := make([]*hit, 0, len(apis))

	for _, api := range apis {
		r := httptest.NewRequest(api.Method, api.Test, nil)
		h.ServeHTTP(w, r)

		hits = append(hits, &hit{
			OK:     w.Body.String() == r.URL.Path,
			Path:   api.Test,
			Want:   api.Brace,
			Actual: w.Body.String(),
		})

		w.Body.Reset()
	}

	return hits
}

// 加载一组 API 到指定的路由中，返回该路由的 http.Handler 接口和所消耗的内存
func loadAPIS(apis []*apis.API, load routers.Load) (http.Handler, uint64) {
	stats := &runtime.MemStats{}

	runtime.GC()
	runtime.ReadMemStats(stats)
	before := stats.HeapAlloc

	h := load(apis)

	runtime.GC()
	runtime.ReadMemStats(stats)
	after := stats.HeapAlloc

	return h, after - before
}
