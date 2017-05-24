// Copyright 2017 by caixw, All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package main

import (
	"net/http"
	"runtime"

	"github.com/caixw/go-http-routing-test/apis"
	"github.com/caixw/go-http-routing-test/routers"
)

func main() {
	//
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
