// Copyright 2017 by caixw, All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package data

type env struct {
	OS   string `json:"os"`
	Arch string `json:"arch"`
	CPU  int    `json:"cpu"`
}

// 每个 API 对应的路由测试数据
type item struct {
	RouterName string `json:"routerName"`     // 路由的名称
	APIName    string `json:"apiName"`        // API 名称
	MemBytes   uint64 `json:"memBytes"`       // 分配的内存
	Bench      *bench `json:"benchmark"`      // 所有的性能数据
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
	Path   string `json:"path"`   // 请求地址
	Want   string `json:"want"`   // 应该匹配的地址
	Actual string `json:"actual"` // 实际访问的地址
}