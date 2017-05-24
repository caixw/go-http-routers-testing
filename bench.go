// Copyright 2017 by caixw, All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/caixw/go-http-routing-test/apis"
)

func bench(apis []*apis.API, h http.Handler, b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		api := apis[i%len(apis)]

		w := httptest.NewRecorder()
		r := httptest.NewRequest(api.Method, api.Test, nil)
		h.ServeHTTP(w, r)

		if w.Body.String() != r.URL.Path {
			b.Errorf("%v: %v:%v", name, w.Body.String(), r.URL.Path)
		}
	}
}
