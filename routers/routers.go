// Copyright 2017 by caixw, All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

// Package routers 各类路由包的定义处。
package routers

import (
	"net/http"

	"github.com/caixw/go-http-routes-testing/apis"
)

var Routers = []*Router{}

type Load func(apis []*apis.API) http.Handler

type Router struct {
	// 路由的名称，必须为一个合法的文件名，且需要与其它路由保持唯一，
	// 建议使用 owner-reop 的形式命名，
	Name string
	URL  string
	Load Load
}

func defaultHandleFunc(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(r.URL.Path))
}
