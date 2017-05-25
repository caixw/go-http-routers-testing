// Copyright 2017 by caixw, All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package routers

import (
	"net/http"

	"github.com/caixw/go-http-routes-testing/apis"
	"github.com/dimfeld/httptreemux"
)

func init() {
	Routers = append(Routers, &Router{
		Name: "dimfeld-httptreemux",
		URL:  "https://github.com/dimfeld/httptreemux",
		Load: dimfeldHTTPTreeMuxLoad,
	})
}

func dimfeldHTTPTreeMuxLoad(apis []*apis.API) http.Handler {
	mux := httptreemux.New()

	for _, api := range apis {
		mux.Handle(api.Method, api.Colon, dimfeldHTTPTreeMuxHandler)
	}

	return mux
}

func dimfeldHTTPTreeMuxHandler(w http.ResponseWriter, r *http.Request, p map[string]string) {
	w.Write([]byte(r.URL.Path))
}
