// Copyright 2017 by caixw, All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"

	"github.com/caixw/go-http-routers-testing/data"
)

const version = "0.1.1+20181224"

const docsDir = "./docs"

func main() {
	v := flag.Bool("v", false, "显示版本号")
	flag.Parse()

	if *v {
		fmt.Println("version:", version)
		return
	}

	data.JSON(filepath.Join(docsDir, "data"), os.Stdout)
}
