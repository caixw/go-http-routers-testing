// Copyright 2017 by caixw, All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package data

type env struct {
	OS   string   `json:"os"`
	Arch string   `json:"arch"`
	CPU  int      `json:"cpu"`
	Go   string   `json:"go"`   // Go 的版本
	Data []string `json:"data"` // 测试数据文件列表
}

type api struct {
	Name    string    `json:"name"`
	Desc    string    `json:"desc"`
	Count   int       `json:"count"`
	Routers []*router `json:"routers"`
}

// 每个 API 对应的路由测试数据
type router struct {
	RouterName        string `json:"name"`     // 路由的名称
	RouterURL         string `json:"url"`      // 路由工具的 URL
	MemBytes          uint64 `json:"memBytes"` // 分配的内存
	NsPerOp           int64  `json:"nsPerOp"`
	AllocsPerOp       int64  `json:"allocsPerOp"`
	AllocedBytesPerOp int64  `json:"allocedBytesPerOp"`
	HitPercent        int    `json:"hitPercent"`     // 命中率
	HitFile           string `json:"hitFile"`        // 保存 hit 记录的文件名
	Hits              []*hit `json:"hits,omitempty"` // 所有的命中数据
}

type hit struct {
	OK     bool   `json:"ok"`     // 是否正确
	Method string `json:"method"` // 请求方法
	Path   string `json:"path"`   // 请求地址
	Want   string `json:"want"`   // 应该匹配的地址
	Body   string `json:"body"`   // 实际输出的内容，默认情况应该就是路由项
}
