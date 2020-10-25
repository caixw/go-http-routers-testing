// SPDX-License-Identifier: MIT

package routers

import (
	"net/http"

	"github.com/caixw/go-http-routers-testing/apis"
	"github.com/gorilla/mux"
)

func init() {
	Routers = append(Routers, &Router{
		Name: "gorilla-mux",
		URL:  "https://github.com/gorilla/mux",
		Load: gorillaMuxLoad,
	})
}

func gorillaMuxLoad(apis []*apis.API) http.Handler {
	mux := mux.NewRouter()

	for _, api := range apis {
		mux.HandleFunc(api.Brace, defaultHandleFunc).Methods(api.Method)
	}

	return mux
}
