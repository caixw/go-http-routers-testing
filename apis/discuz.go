// Copyright 2017 by caixw, All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package apis

import "net/http"

func init() {
	discuz.init()

	APIS = append(APIS, discuz)
}

var discuz = &Collection{
	Name: "Discuz Routes",
	Desc: "DZ 风格的路由定义",
	APIS: []*API{
		{Method: http.MethodGet, Brace: "/forum-{fid}-{page}.html"},
		{Method: http.MethodGet, Brace: "/thread-{tid}-{page}-{ordertype}.html"},
	},
}
