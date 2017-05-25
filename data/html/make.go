// Copyright 2017 by caixw, All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package main

import (
	"bytes"
	"go/format"
	"io/ioutil"
	"os"
)

const (
	destFile = "./static.go" // 产生的代码文件
	pkgName  = "html"
)

var srcFiles = map[string]string{
	"Index": "./index.html",
	"Hit":   "./hit.html",
}

func main() {

	dest, err := os.Create(destFile)
	if err != nil {
		panic(err)
	}
	defer dest.Close()

	buf := new(bytes.Buffer)
	buf.WriteString("// 这是自动产生的文件，不需要修改")
	buf.WriteString("\n\n")

	buf.WriteString("package ")
	buf.WriteString(pkgName)
	buf.WriteString("\n\n")

	for name, path := range srcFiles {
		f, err := os.Open(path)
		if err != nil {
			panic(err)
		}

		d, err := ioutil.ReadAll(f)
		if err != nil {
			panic(err)
		}

		buf.WriteString("var ")
		buf.WriteString(name)
		buf.WriteString("=[]byte(`")
		buf.Write(d)
		buf.WriteString("`)\n")
	}

	// 格式化
	bs, err := format.Source(buf.Bytes())
	if err != nil {
		panic(err)
	}

	_, err = dest.Write(bs)
	if err != nil {
		panic(err)
	}
}
