// Copyright 2017 by caixw, All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package main

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

// 数据存放目录
const dataDir = "./docs/data"

// 每个 API 对应的路由测试数据
type router struct {
	RouterName string `json:"routerName"` // 路由的名称
	APIName    string `json:"apiName"`    // API 名称
	MemBytes   uint64 `json:"memBytes"`   // 分配的内存
	Bench      *bench `json:"benchmark"`  // 所有的性能数据
	Hits       []*hit `json:"hits"`       // 所有的命中数据
}

type bench struct {
	NsPerOp           int64 `json:"nsPerOp"`
	AllocsPerOp       int64 `json:"allocsPerOp"`
	AllocedBytesPerOp int64 `json:"allocedBytesPerOp"`
}

type hit struct {
	OK     bool   `json:"ok"`     // 是否正确
	Path   string `json:"path"`   // 请求地址
	Want   string `json:"want"`   // 应该匹配的地址
	Actual string `json:"actual"` // 实际访问的地址
}

func main() {
	for index, c := range apis.APIS {
		rs := make([]*router, 0, len(routers.Routers))

		for _, r := range routers.Routers {
			rs = append(rs, single(c, r))
		}

		bs, err := json.MarshalIndent(rs, "", "  ")
		if err != nil {
			panic(err)
		}

		path := filepath.Join(dataDir, strconv.Itoa(index)+".json")
		f, err := os.Create(path)
		if err != nil {
			panic(err)
		}
		defer f.Close()
		f.Write(bs)
	}
}

func single(c *apis.Collection, r *routers.Router) *router {
	ret := &router{
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
