// Copyright 2017 by caixw, All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package main

import (
	"os"
	"path/filepath"

	"github.com/caixw/go-http-routers-testing/data"
)

const docsDir = "./docs"

func main() {
	data.JSON(filepath.Join(docsDir, "data"), os.Stdout)
}
