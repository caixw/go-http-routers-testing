// SPDX-FileCopyrightText: 2017-2024 caixw
//
// SPDX-License-Identifier: MIT

package apis

import "net/http"

func init() {
	discuz.init()
}

var discuz = &Collection{
	ID:   "discuz",
	Name: "Discuz Routes",
	Desc: "DZ 风格的路由定义",
	APIS: []*API{
		{Method: http.MethodGet, Brace: "/forum-{fid}-{page}.html"},
		{Method: http.MethodGet, Brace: "/thread-{tid}-{page}-{ordertype}.html"},
	},
}
