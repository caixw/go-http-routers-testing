// SPDX-License-Identifier: MIT

package routers

import (
	"net/http"

	"github.com/gorilla/mux"

	"github.com/caixw/go-http-routers-testing/apis"
)

func init() {
	Routers = append(Routers, &Router{
		Name: "gorilla-mux/v1.8",
		URL:  "https://github.com/gorilla/mux",
		Load: gorillaMuxLoad,
	})
}

func gorillaMuxLoad(apis []*apis.API) ServeFunc {
	router := mux.NewRouter()
	for _, api := range apis {
		router.HandleFunc(api.Brace, func(w http.ResponseWriter, r *http.Request) {
			_ = mux.Vars(r)
			w.Write([]byte(r.URL.Path))
		}).Methods(api.Method)
	}
	return stdServeFunc(router)
}
