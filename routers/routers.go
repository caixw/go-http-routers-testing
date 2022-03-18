// Copyright 2017 by caixw, All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

// Package routers 各类路由包的定义处
package routers

import (
	"net/http"
	"net/http/httptest"
	"sort"

	"github.com/issue9/sliceutil"

	"github.com/caixw/go-http-routers-testing/apis"
)

var routers = make([]*Router, 0, 10)

// Load 加载路由的函数
type Load func(apis []*apis.API) ServeFunc

// ServeFunc 传入一个 api，返回处理之后的内容。
type ServeFunc func(*apis.API) string

// Router 路由的相关信息
type Router struct {
	ID   string // 唯一 ID，建议使用 owner-repo-version 的形式命名。
	Name string // 路由的名称，建议使用 owner-repo/version 的形式命名。
	URL  string
	Load Load
}

func stdServeFunc(h http.Handler) ServeFunc {
	return func(api *apis.API) string {
		w := httptest.NewRecorder()
		r, err := http.NewRequest(api.Method, api.Test, nil)
		if err != nil {
			panic(err)
		}
		h.ServeHTTP(w, r)
		return w.Body.String()
	}
}

func Routers() []*Router {
	dup := sliceutil.Dup(routers, func(i, j *Router) bool { return i.ID == j.ID })
	if len(dup) > 0 {
		panic("存在重复的 ID 值" + routers[dup[0]].ID)
	}

	sort.SliceStable(routers, func(i, j int) bool {
		return routers[i].ID < routers[j].ID
	})

	return routers
}
