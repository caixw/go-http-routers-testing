// SPDX-FileCopyrightText: 2017-2024 caixw
//
// SPDX-License-Identifier: MIT

package main

import (
	"log"

	"github.com/caixw/go-http-routers-testing/data"
)

const docsDir = "./docs"

func main() {
	if err := data.JSON(docsDir, log.Default()); err != nil {
		panic(err)
	}
}
