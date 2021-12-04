// Copyright 2017 by caixw, All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package routers

import (
	"net/http"

	"github.com/issue9/mux/v5"

	"github.com/caixw/go-http-routers-testing/apis"
)

func init() {
	Routers = append(Routers, &Router{
		Name: "issue9-mux/v5",
		URL:  "https://github.com/issue9/mux",
		Load: issue9MuxLoad,
	})
}

func issue9MuxLoad(apis []*apis.API) ServeFunc {
	router := mux.NewRouter("")
	for _, api := range apis {
		router.HandleFunc(api.Brace, func(w http.ResponseWriter, r *http.Request) {
			_ = mux.GetParams(r)
			w.Write([]byte(r.URL.Path))
		}, api.Method)
	}
	return stdServeFunc(router)
}
