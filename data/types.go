// Copyright 2017 by caixw, All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package data

type env struct {
	OS   string   `json:"os"`
	Arch string   `json:"arch"`
	CPU  int      `json:"cpu"`
	Data []string `json:"data"` // 测试数据文件列表
}

// 每个 API 对应的路由测试数据
type item struct {
	RouterName string `json:"routerName"`     // 路由的名称
	RouterURL  string `json:"routerURL"`      // 路由工具的 URL
	APIName    string `json:"apiName"`        // API 名称
	APICount   int    `json:"apiCount"`       // 本次测试的 API 数量
	MemBytes   uint64 `json:"memBytes"`       // 分配的内存
	Bench      *bench `json:"bench"`          // 所有的性能数据
	HitPrecent int    `json:"hitPrecent"`     // 命中率
	HitFile    string `json:"hitFile"`        // 保存 hit 记录的文件名
	Hits       []*hit `json:"hits,omitempty"` // 所有的命中数据
}

type bench struct {
	NsPerOp           int64 `json:"nsPerOp"`
	AllocsPerOp       int64 `json:"allocsPerOp"`
	AllocedBytesPerOp int64 `json:"allocedBytesPerOp"`
}

type hit struct {
	OK     bool   `json:"ok"`     // 是否正确
	Method string `json:"method"` // 请求方法
	Path   string `json:"path"`   // 请求地址
	Want   string `json:"want"`   // 应该匹配的地址
	Body   string `json:"body"`   // 实际输出的内容，默认情况应该就是路由项
}
