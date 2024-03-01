// SPDX-FileCopyrightText: 2017-2024 caixw
//
// SPDX-License-Identifier: MIT

//go:build ignore

package main

import (
	"flag"
	"fmt"
	"net/http"
)

const docsDir = "./docs"

func main() {
	p := flag.String("p", ":8080", "运行本地服务的端口")
	flag.Parse()
	if *p != "" {
		if (*p)[0] != ':' {
			*p = ":" + *p
		}

		http.Handle("/", http.StripPrefix("/", http.FileServer(http.Dir(docsDir))))
		fmt.Printf("运行服务: %s\n", *p)

		err := http.ListenAndServe(*p, nil)
		if err != nil && err != http.ErrServerClosed {
			panic(err)
		}
		return
	}
}
