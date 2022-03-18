// Copyright 2017 by caixw, All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package routers

import (
	"net/http"

	"github.com/issue9/mux/v6"

	"github.com/caixw/go-http-routers-testing/apis"
)

func init() {
	Routers = append(Routers, &Router{
		ID:   "issue9-mux-v6",
		Name: "issue9-mux/v6",
		URL:  "https://github.com/issue9/mux",
		Load: issue9MuxLoad,
	})
}

func issue9MuxLoad(apis []*apis.API) ServeFunc {
	router := mux.NewRouter("", nil)
	for _, api := range apis {
		router.Handle(api.Brace, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			_ = mux.GetParams(r)
			w.Write([]byte(r.URL.Path))
		}), api.Method)
	}
	return stdServeFunc(router)
}
