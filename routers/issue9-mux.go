// SPDX-FileCopyrightText: 2017-2024 caixw
//
// SPDX-License-Identifier: MIT

package routers

import (
	"net/http"

	"github.com/issue9/mux/v7/examples/std"

	"github.com/caixw/go-http-routers-testing/apis"
)

func init() {
	routers = append(routers, &Router{
		ID:   "issue9-mux-v7",
		Name: "issue9-mux/v7",
		URL:  "https://github.com/issue9/mux",
		Load: issue9MuxLoad,
	})
}

func issue9MuxLoad(c *apis.Collection) ServeFunc {
	router := std.NewRouter("")
	for _, api := range c.APIS {
		router.Handle(api.Brace, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			_ = std.GetParams(r)
			w.Write([]byte(r.URL.Path))
		}), api.Method)
	}
	return stdServeFunc(router)
}
