// Copyright 2017 by caixw, All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

// Package data 处理数据
package data

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"runtime"
	"sort"
	"strconv"
	"testing"

	"github.com/caixw/go-http-routers-testing/apis"
	"github.com/caixw/go-http-routers-testing/routers"
)

// JSON 输出 JSON 数据，数据按路由进行分组。
func JSON(dir string, log io.Writer) error {
	sort.SliceStable(apis.APIS, func(i, j int) bool {
		return apis.APIS[i].Name > apis.APIS[j].Name
	})

	env := getEnv()
	env.Data = make([]string, 0, len(apis.APIS))

	for cindex, c := range apis.APIS {
		cdata := make([]*item, 0, len(routers.Routers))

		fmt.Fprintf(log, "开始测试 %v \n", c.Name)

		for rindex, r := range routers.Routers {
			fmt.Fprint(log, "    ", r.Name, "......")

			filename := strconv.Itoa(cindex) + "-" + strconv.Itoa(rindex) + ".json"
			data := single(c, r)
			data.HitFile = filename
			if err := writeJSON(filepath.Join(dir, filename), data.Hits); err != nil {
				return err
			}

			data.Hits = nil
			cdata = append(cdata, data)

			fmt.Fprintln(log, "[OK]")
		}

		filename := strconv.Itoa(cindex) + ".json"
		env.Data = append(env.Data, filename)
		if err := writeJSON(filepath.Join(dir, filename), cdata); err != nil {
			return err
		}

		fmt.Fprintf(log, "完成 %v 测试\n\n", c.Name)
	} // end for routers.Routers

	return writeJSON(filepath.Join(dir, "env.json"), env)
}

func writeJSON(path string, obj interface{}) error {
	bs, err := json.MarshalIndent(obj, "", "  ")
	if err != nil {
		return err
	}

	f, err := os.Create(path)
	if err != nil {
		return err
	}
	defer f.Close()

	_, err = f.Write(bs)
	return err
}

func getEnv() *env {
	return &env{
		OS:   runtime.GOOS,
		Arch: runtime.GOARCH,
		CPU:  runtime.NumCPU(),
		Go:   runtime.Version(),
	}
}

func single(c *apis.Collection, r *routers.Router) *item {
	ret := &item{
		RouterName: r.Name,
		RouterURL:  r.URL,
		APIName:    c.Name,
		APIDesc:    c.Desc,
		APICount:   len(c.APIS),
	}

	// 加载
	h, size := loadAPIS(c.APIS, r.Load)
	ret.MemBytes = size

	// 命中情况
	ret.Hits = testHit(c.APIS, h)
	var cnt int
	for _, hit := range ret.Hits {
		if hit.OK {
			cnt++
		}
	}
	ret.HitPercent = cnt * 100 / len(ret.Hits)

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
			Method: r.Method,
			Path:   api.Test,
			Want:   api.Brace,
			Body:   w.Body.String(),
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

	runtime.ReadMemStats(stats)
	after := stats.HeapAlloc

	return h, after - before
}
