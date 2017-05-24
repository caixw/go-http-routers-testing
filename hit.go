// Copyright 2017 by caixw, All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package main

import (
	"net/http"
	"net/http/httptest"

	"github.com/caixw/go-http-routing-test/apis"
)

func hit(apis []*apis.API, h http.Handler) bool {
	w := httptest.NewRecorder()

	for _, api := range apis {
		r := httptest.NewRequest(api.Method, api.Test, nil)
		h.ServeHTTP(w, r)

		if w.Body.String() != r.URL.Path {
			// TODO
			b.Errorf("%v: %v:%v", name, w.Body.String(), r.URL.Path)
		}
	}
}
