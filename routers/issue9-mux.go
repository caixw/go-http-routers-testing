// Copyright 2017 by caixw, All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package routers

import (
	"net/http"

	"github.com/issue9/mux/v4"
	"github.com/issue9/mux/v4/group"

	"github.com/caixw/go-http-routers-testing/apis"
)

func init() {
	Routers = append(Routers, &Router{
		Name: "issue9-mux",
		URL:  "https://github.com/issue9/mux",
		Load: issue9MuxLoad,
	})
}

func issue9MuxLoad(apis []*apis.API) http.Handler {
	mux := mux.Default()
	r, ok := mux.NewRouter("any", group.MatcherFunc(group.Any))
	if !ok {
		panic("初始化 issue9-mux 出错")
	}

	for _, api := range apis {
		err := r.HandleFunc(api.Brace, defaultHandleFunc, api.Method)
		if err != nil {
			panic(err)
		}
	}

	return mux
}
