// SPDX-FileCopyrightText: 2017-2024 caixw
//
// SPDX-License-Identifier: MIT

package routers

import (
	"net/http"

	"github.com/dimfeld/httptreemux/v5"

	"github.com/caixw/go-http-routers-testing/apis"
)

func init() {
	routers = append(routers, &Router{
		ID:   "dimfeld-httptreemux-v5",
		Name: "dimfeld-httptreemux/v5",
		URL:  "https://github.com/dimfeld/httptreemux",
		Load: dimfeldHTTPTreeMuxLoad,
	})
}

func dimfeldHTTPTreeMuxLoad(c *apis.Collection) ServeFunc {
	mux := httptreemux.New()
	for _, api := range c.APIS {
		mux.Handle(api.Method, api.Colon, dimfeldHTTPTreeMuxHandler)
	}
	return stdServeFunc(mux)
}

func dimfeldHTTPTreeMuxHandler(w http.ResponseWriter, r *http.Request, p map[string]string) {
	w.Write([]byte(r.URL.Path))
}
