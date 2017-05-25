// Copyright 2017 by caixw, All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package main

import (
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/caixw/go-http-routes-testing/data"
)

// 数据存放目录
const (
	docsDir = "./docs"
	dataDir = "./docs/data"
)

const version = "0.1.1+20170525"

// 默认的输出通首
var output = os.Stdout

func main() {
	typ := flag.String("type", "html", "指定输出的数据类型，可以是 HTML 或是 JSON")
	v := flag.Bool("v", false, "显示版本号")

	if *v {
		fmt.Println("version:", version)
		return
	}

	switch strings.ToLower(*typ) {
	case "html":
		data.HTML(docsDir, output)
	case "json":
		data.JSON(dataDir, output)
	default:
		panic("无效的 type 值")
	}
}
