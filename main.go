// Copyright 2017 by caixw, All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package main

import (
	"os"

	"github.com/caixw/go-http-routes-testing/data"
)

// 数据存放目录
const dataDir = "./docs/data"

func main() {
	data.JSON(dataDir, os.Stdout)
}
