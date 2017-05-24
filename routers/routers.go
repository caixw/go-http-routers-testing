// Copyright 2017 by caixw, All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package routers

import (
	"net/http"

	"github.com/caixw/go-http-routing-test/apis"
)

var Routers = []*Router{}

type Load func(apis []*apis.API) http.Handler

type Router struct {
	Name string
	URL  string
	Load Load
}

func defaultHandleFunc(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(r.URL.Path))
}
