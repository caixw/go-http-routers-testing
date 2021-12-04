// Copyright 2017 by caixw, All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

// Package routers 各类路由包的定义处
package routers

import (
	"net/http"
	"net/http/httptest"

	"github.com/caixw/go-http-routers-testing/apis"
)

// Routers 所有的路由
var Routers = make([]*Router, 0, 10)

// Load 加载路由的函数
type Load func(apis []*apis.API) ServeFunc

// ServeFunc 传入一个 api，返回处理之后的内容。
type ServeFunc func(*apis.API) string

// Router 路由的相关信息
type Router struct {
	Name string // 路由的名称，建议使用 owner-repo 的形式命名。
	URL  string
	Load Load
}

func stdServeFunc(h http.Handler) ServeFunc {
	return func(api *apis.API) string {
		w := httptest.NewRecorder()
		r, _ := http.NewRequest(api.Method, api.Test, nil)
		h.ServeHTTP(w, r)
		return w.Body.String()
	}
}
