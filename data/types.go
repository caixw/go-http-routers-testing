// Copyright 2017 by caixw, All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package data

type env struct {
	OS       string   `json:"os"`
	Arch     string   `json:"arch"`
	CPU      int      `json:"cpu"`
	Go       string   `json:"go"`       // Go 的版本
	APIFiles []string `json:"apiFiles"` // 测试数据文件列表
	APIs     []*api   `json:"-"`        // 仅用于 html 输出
}

type api struct {
	ID      string    `json:"id"`
	Name    string    `json:"name"`
	Desc    string    `json:"desc"`
	Count   int       `json:"count"`
	Routers []*router `json:"routers"`
	File    string    `json:"-"` // 仅用于 html 输出
	Min     *testData `json:"-"` // 记录最小的值，仅用于 html 输出
}

type testData struct {
	MemBytes          uint64
	NsPerOp           int64
	AllocsPerOp       int64
	AllocedBytesPerOp int64
}

// 每个 API 对应的路由测试数据
type router struct {
	ID                string `json:"id"`
	Name              string `json:"name"`     // 路由的名称
	URL               string `json:"url"`      // 路由工具的 URL
	MemBytes          uint64 `json:"memBytes"` // 分配的内存
	NsPerOp           int64  `json:"nsPerOp"`
	AllocsPerOp       int64  `json:"allocsPerOp"`
	AllocedBytesPerOp int64  `json:"allocedBytesPerOp"`
	MissFile          string `json:"missFile,omitempty"` // 保存未命中记录的文件名
	Misses            int    `json:"misses,omitempty"`   // 未命中的数量
	missesData        []*miss
}

type miss struct {
	Method string `json:"method"` // 请求方法
	Path   string `json:"path"`   // 请求地址
	Want   string `json:"want"`   // 应该匹配的地址
	Body   string `json:"body"`   // 实际输出的内容，默认情况应该就是路由项
}
