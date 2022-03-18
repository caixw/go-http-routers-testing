// Copyright 2017 by caixw, All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

// Package apis 预定义的各类型的 API 数据
package apis

import (
	"sort"
	"strings"

	"github.com/issue9/sliceutil"
)

var apis = []*Collection{
	{
		ID:   "all",
		Name: "所有",
		Desc: "所有接口混合测试",
		APIS: make([]*API, 0, 1000),
	},
}

// Collection 表示一组 API 的定义。
type Collection struct {
	ID   string
	Name string
	Desc string
	APIS []*API
}

// API 单个 API 接口的定义
type API struct {
	Method string // 请求方法
	Brace  string // {id} 风格的路由项
	Colon  string // :id 风格的路由项
	Test   string // 本条路由对应的测试地址
}

func (c *Collection) init() {
	for _, api := range c.APIS { // 初始化另外两个字段
		path := strings.Replace(api.Brace, "}", "", -1)
		api.Test = strings.Replace(path, "{", "", -1)
		api.Colon = strings.Replace(path, "{", ":", -1)
	}

	apis = append(apis, c)
	apis[0].APIS = append(apis[0].APIS, c.APIS...)

}

func APIs() []*Collection {
	dup := sliceutil.Dup(apis, func(i, j *Collection) bool { return i.ID == j.ID })
	if len(dup) > 0 {
		panic("存在重复的 ID 值" + apis[dup[0]].ID)
	}

	sort.SliceStable(apis, func(i, j int) bool {
		return apis[i].ID > apis[j].ID
	})

	return apis
}
