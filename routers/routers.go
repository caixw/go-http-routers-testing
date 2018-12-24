// Copyright 2017 by caixw, All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

// Package routers 各类路由包的定义处。
package routers

import (
	"net/http"

	"github.com/caixw/go-http-routers-testing/apis"
)

// Routers 所有的路由
var Routers = []*Router{}

// Load 加载路由的函数
type Load func(apis []*apis.API) http.Handler

// Router 路由的相关信息
type Router struct {
	// 路由的名称，建议使用 owner-reop 的形式命名，
	Name string
	URL  string
	Load Load
}

func defaultHandleFunc(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(r.URL.Path))
}
