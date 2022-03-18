// Copyright 2017 by caixw, All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

// Package data 处理数据
package data

import (
	"net/http"
	"net/http/httptest"
	"runtime"
	"runtime/metrics"
	"testing"

	"github.com/caixw/go-http-routers-testing/apis"
	"github.com/caixw/go-http-routers-testing/routers"
)

const dataDirName = "data"

func getEnv() *env {
	return &env{
		OS:   runtime.GOOS,
		Arch: runtime.GOARCH,
		CPU:  runtime.NumCPU(),
		Go:   runtime.Version(),
	}
}

func testRouter(c *apis.Collection, r *routers.Router) *router {
	ret := &router{
		ID:   r.ID,
		Name: r.Name,
		URL:  r.URL,
	}

	// 加载
	s, size := loadAPIS(c.APIS, r.Load)
	ret.MemBytes = size

	// bench
	rslt := testing.Benchmark(func(b *testing.B) {
		b.ReportAllocs()
		b.ResetTimer()

		for i := 0; i < b.N; i++ {
			api := c.APIS[i%len(c.APIS)]
			s(api)
		}
	})
	ret.NsPerOp = rslt.NsPerOp()
	ret.AllocsPerOp = rslt.AllocsPerOp()
	ret.AllocedBytesPerOp = rslt.AllocedBytesPerOp()

	// 命中情况
	ret.missesData = testMiss(c.APIS, s)
	ret.Misses = len(ret.missesData)

	return ret
}

func testMiss(apis []*apis.API, s routers.ServeFunc) []*miss {
	w := httptest.NewRecorder()
	misses := make([]*miss, 0, len(apis))

	for _, api := range apis {
		r, err := http.NewRequest(api.Method, api.Test, nil)
		if err != nil {
			panic(err)
		}

		body := s(api)
		if body != r.URL.Path {
			misses = append(misses, &miss{
				Method: r.Method,
				Path:   api.Test,
				Want:   api.Brace,
				Body:   body,
			})
		}

		w.Body.Reset()
	}

	return misses
}

// 加载一组 API 到指定的路由中，返回该路由的处理方法和所消耗的内存。
func loadAPIS(apis []*apis.API, load routers.Load) (routers.ServeFunc, uint64) {
	sample := make([]metrics.Sample, 1)
	sample[0].Name = "/gc/heap/allocs:bytes"

	runtime.GC()
	metrics.Read(sample)
	before := sample[0].Value

	s := load(apis)

	metrics.Read(sample)
	after := sample[0].Value

	return s, after.Uint64() - before.Uint64()
}
