// Copyright 2017 by caixw, All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

// Package apis 预定义的各类型的 API 数据。
package apis

import "strings"

// APIS 所有 API 的集体
var APIS = []*Collection{
	{
		Name: "所有",
		Desc: "所有接口混合测试",
		APIS: make([]*API, 0, 1000),
	},
}

// Collection 表示一组 API 的定义。
type Collection struct {
	Name string
	Desc string
	APIS []*API
}

func addCollection(c *Collection) {
	APIS = append(APIS, c)
	APIS[0].APIS = append(APIS[0].APIS, c.APIS...)
}

// API 单个 API 接口的定义
type API struct {
	Method string // 请求方法
	Brace  string // {id} 风格的路由项
	Colon  string // :id 风格的路由项
	Test   string // 本条路由对应的测试地址
}

// 初始化另外两个字段
func (api *API) init() {
	path := strings.Replace(api.Brace, "}", "", -1)

	api.Test = strings.Replace(path, "{", "", -1)
	api.Colon = strings.Replace(path, "{", ":", -1)
}

func (c *Collection) init() {
	for _, api := range c.APIS {
		api.init()
	}

	addCollection(c)
}
