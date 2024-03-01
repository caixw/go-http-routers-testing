// SPDX-FileCopyrightText: 2017-2024 caixw
//
// SPDX-License-Identifier: MIT

package routers

import (
	"net/http"

	"github.com/gorilla/mux"

	"github.com/caixw/go-http-routers-testing/apis"
)

func init() {
	routers = append(routers, &Router{
		ID:   "gorilla-mux",
		Name: "gorilla-mux",
		URL:  "https://github.com/gorilla/mux",
		Load: gorillaMuxLoad,
	})
}

func gorillaMuxLoad(c *apis.Collection) ServeFunc {
	router := mux.NewRouter()
	for _, api := range c.APIS {
		router.HandleFunc(api.Brace, func(w http.ResponseWriter, r *http.Request) {
			_ = mux.Vars(r)
			w.Write([]byte(r.URL.Path))
		}).Methods(api.Method)
	}
	return stdServeFunc(router)
}
