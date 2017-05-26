// Copyright 2017 by caixw, All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package routers

import (
	"net/http"

	"github.com/caixw/go-http-routes-testing/apis"
	"github.com/issue9/mux"
)

func init() {
	Routers = append(Routers, &Router{
		Name: "issue9-mux",
		URL:  "https://github.com/issue9/mux",
		Load: issue9MuxLoad,
	})
}

func issue9MuxLoad(apis []*apis.API) http.Handler {
	mux := mux.New(false, false, nil, nil)

	for _, api := range apis {
		err := mux.HandleFunc(api.Brace, defaultHandleFunc, api.Method)
		if err != nil {
			panic(err)
		}
	}

	return mux
}
