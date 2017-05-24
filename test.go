// Copyright 2017 by caixw, All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package main

import (
	"net/http"
	"net/http/httptest"
	"runtime"
	"testing"

	"github.com/caixw/go-http-routing-test/apis"
	"github.com/caixw/go-http-routing-test/routers"
)

// 每个 API 的测试数据
type api struct {
	Name       string `json:"name"`       // API 组的名称
	HitPrecent int    `json:"hitPrecent"` // 命中率
	MemBytes   uint64 `json:"memBytes"`   // 分配的内存
	Bench      *bench `json:"benchmark"`  // 所有的性能数据

	Hits []*hit `json:"hits"` // 所有的命中数据
}

type bench struct {
	NsPerOp      int64 `json:"nsPerOp"`
	AllocsPerOp  int64 `json:"allocsPerOp"`
	AllocedPerOp int64 `json:"allocedPerOp"`
}

type hit struct {
	OK     bool   `json:"ok"`     // 是否正确
	Path   string `json:"path"`   // 请求地址
	Want   string `json:"want"`   // 应该匹配的地址
	Actual string `json:"actual"` // 实际访问的地址
}

func testHit(apis []*apis.API, h http.Handler) bool {
	w := httptest.NewRecorder()

	for _, api := range apis {
		r := httptest.NewRequest(api.Method, api.Test, nil)
		h.ServeHTTP(w, r)

		if w.Body.String() != r.URL.Path {
			// TODO
			b.Errorf("%v: %v:%v", name, w.Body.String(), r.URL.Path)
		}
	}
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

func testBench(apis []*apis.API, h http.Handler, b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		api := apis[i%len(apis)]

		w := httptest.NewRecorder()
		r := httptest.NewRequest(api.Method, api.Test, nil)
		h.ServeHTTP(w, r)

		if w.Body.String() != r.URL.Path {
			b.Errorf("%v: %v:%v", name, w.Body.String(), r.URL.Path)
		}
	}
}
