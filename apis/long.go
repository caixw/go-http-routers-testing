// SPDX-FileCopyrightText: 2017-2024 caixw
//
// SPDX-License-Identifier: MIT

package apis

import "net/http"

func init() {
	long.init()
}

var long = &Collection{
	ID:   "long",
	Name: "Long Routes",
	Desc: "路径很长的路由项",
	APIS: []*API{
		{Method: http.MethodGet, Brace: "/loooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooog/{p1}/loooooooooooooooooooooooooooooooooooooooogooooooooooooooooooooooooooooooog/{p2}.html"},
		{Method: http.MethodPut, Brace: "/loooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooogooooooooooooooooooooooooooooooooooooog/{p1}/loooooooooooooooooooooooooooooooooooooooogooooooooooooooooooooooooooooooog/{p2}.html"},
		{Method: http.MethodPut, Brace: "/looooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooog/{p1}/loooooooooooooooooooooooooooooooooooooooogooooooooooooooooooooooooooooooog/{p2}.html"},
		{Method: http.MethodPut, Brace: "/looooooooooooooooooooooooooooooooooogooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooog/{p1}/looooooooooooooogoooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooogooooooooooooooooooooooooooooooog/{p2}.html"},
		{Method: http.MethodPut, Brace: "/looooooooooooooooooooooooooooooooooogooooooooooooooogooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooog/{p1}/loooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooogoooooooooooooooooooooooogooooooooooooooooooooooooooooooog/{p2}.html"},
		{Method: http.MethodGet, Brace: "/looooooooooooooooooooooooooooooooooogooooooooooooooogooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooog/{p1}/looooooooooooooogoooooooooooooooooooooooogooooooooooooooooooooooooooooooooooooooooooooooooooooooooooog/{p2}.html"},
		{Method: http.MethodGet, Brace: "/looooooooooooooooooooooooooooooooooogooooooooooooooogoooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooog/{p1}/looooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooogoooooooooooogooooooooooogooooooooooooooooooooooooooooooog/{p2}.html"},
		{Method: http.MethodGet, Brace: "/loooooooooooogoooooooooooooooooooooogooooooooooooooogoooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooog/{p1}/looooooooooooooogoooooooooooogooooooooooogooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooog/{p2}.html"},
		{Method: http.MethodPut, Brace: "/loooooooooooogoooooooooooooooooooooogooooooooooooooogooooooooooooooooooooooooooooooooooooooooooooooooooooooooog/{p1}/looooooooooooooogoooooooooooogooooooooooogoooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooog/{p2}.html"},
		{Method: http.MethodPost, Brace: "/loooooooooooogoooooooooooooooooooooogooooooooooooooogooooooooooooooooooooooooooooooooooooooooooog/{p1}/looooooooooooooogoooooooooooogooooooooooogooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooog/{p2}.html"},
		{Method: http.MethodPost, Brace: "/loooooooooooogoooooooooooooooooooooogooooooooooooooogoooooooooooooooooooooooooooooooooog/{p1}/looooooooooooooogoooooooooooogooooooooooogoooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooog.html"},
		{Method: http.MethodPost, Brace: "/loooooooooooogoooooooooooooooooooooooooooooooooooooooooooooogoooooogoooooooogoooooooooooooooooooog/{p1}/looooooooooooooogoooooooooooogooooooooooogoooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooog.html"},
		{Method: http.MethodPost, Brace: "/loooooooooooogooooooooooooooooooooooooooooooooogoooooogoooooooogoooooooooooooooooooog/{p1}/loooooogoooooooogoooooooooooogooooooooooogooooooooooooooooooooooooooooooooooooooooooooooooooog.html"},
		{Method: http.MethodPost, Brace: "/aoooooooooooogooooooooooooooooooooooooooooooooogoooooogoooooooogoooooooooooooooooooog/{p1}/loooooogoooooooogoooooooooooogooooooooooogooooooooooooooooooooooooooooooooooooooooooooooooooog.html"},
		{Method: http.MethodPost, Brace: "/boooooooooooogooooooooooooooooooooooooooooooooogoooooogoooooooogoooooooooooooooooooog/{p1}/loooooogoooooooogoooooooooooogooooooooooogooooooooooooooooooooooooooooooooooooooooooooooooooog.html"},
		{Method: http.MethodPost, Brace: "/coooooooooooogooooooooooooooooooooooooooooooooogoooooogoooooooogoooooooooooooooooooog/{p1}/loooooogoooooooogoooooooooooogooooooooooogooooooooooooooooooooooooooooooooooooooooooooooooooog.html"},
	},
}
