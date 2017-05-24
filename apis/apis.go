// Copyright 2017 by caixw, All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package apis

// API 对一 API 的定义
type API struct {
	Method string // 请求方法
	Brace  string // {id} 风格的路由项
	Colon  string // :id 风格的路由项
	Test   string // 本条路由对应的测试地址
}

type APIS struct {
	Name string
	APIS []*API
}
