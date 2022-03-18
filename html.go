// SPDX-License-Identifier: MIT

//go:build ignore

package main

import (
	"log"

	"github.com/caixw/go-http-routers-testing/data"
)

const docsDir = "./docs"

func main() {
	if err := data.HTML(docsDir, log.Default()); err != nil {
		panic(err)
	}
}
