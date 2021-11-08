// SPDX-License-Identifier: MIT

package apis

import "net/http"

func init() {
	params2.init()
}

var params2 = &Collection{
	Name: "2 Params",
	Desc: "两个参数的路由",
	APIS: []*API{
		{Method: http.MethodGet, Brace: "/v0/{p1}/0/{p2}.html"},
		{Method: http.MethodGet, Brace: "/v0/{p1}/1/{p2}.html"},
		{Method: http.MethodGet, Brace: "/v0/{p1}/2/{p2}.html"},
		{Method: http.MethodGet, Brace: "/v0/{p1}/3/{p2}.html"},
		{Method: http.MethodGet, Brace: "/v0/{p1}/4/{p2}.html"},
		{Method: http.MethodGet, Brace: "/v0/{p1}/5/{p2}.html"},
		{Method: http.MethodGet, Brace: "/v0/{p1}/6/{p2}.html"},
		{Method: http.MethodGet, Brace: "/v0/{p1}/7/{p2}.html"},
		{Method: http.MethodGet, Brace: "/v0/{p1}/8/{p2}.html"},
		{Method: http.MethodGet, Brace: "/v0/{p1}/9/{p2}.html"},

		{Method: http.MethodGet, Brace: "/v1/{p1}/0/{p2}.html"},
		{Method: http.MethodGet, Brace: "/v1/{p1}/1/{p2}.html"},
		{Method: http.MethodGet, Brace: "/v1/{p1}/2/{p2}.html"},
		{Method: http.MethodGet, Brace: "/v1/{p1}/3/{p2}.html"},
		{Method: http.MethodGet, Brace: "/v1/{p1}/4/{p2}.html"},
		{Method: http.MethodGet, Brace: "/v1/{p1}/5/{p2}.html"},
		{Method: http.MethodGet, Brace: "/v1/{p1}/6/{p2}.html"},
		{Method: http.MethodGet, Brace: "/v1/{p1}/7/{p2}.html"},
		{Method: http.MethodGet, Brace: "/v1/{p1}/8/{p2}.html"},
		{Method: http.MethodGet, Brace: "/v1/{p1}/9/{p2}.html"},

		{Method: http.MethodGet, Brace: "/v2/{p1}/0/{p2}.html"},
		{Method: http.MethodGet, Brace: "/v2/{p1}/1/{p2}.html"},
		{Method: http.MethodGet, Brace: "/v2/{p1}/2/{p2}.html"},
		{Method: http.MethodGet, Brace: "/v2/{p1}/3/{p2}.html"},
		{Method: http.MethodGet, Brace: "/v2/{p1}/4/{p2}.html"},
		{Method: http.MethodGet, Brace: "/v2/{p1}/5/{p2}.html"},
		{Method: http.MethodGet, Brace: "/v2/{p1}/6/{p2}.html"},
		{Method: http.MethodGet, Brace: "/v2/{p1}/7/{p2}.html"},
		{Method: http.MethodGet, Brace: "/v2/{p1}/8/{p2}.html"},
		{Method: http.MethodGet, Brace: "/v2/{p1}/9/{p2}.html"},

		{Method: http.MethodGet, Brace: "/v3/{p1}/0/{p2}.html"},
		{Method: http.MethodGet, Brace: "/v3/{p1}/1/{p2}.html"},
		{Method: http.MethodGet, Brace: "/v3/{p1}/2/{p2}.html"},
		{Method: http.MethodGet, Brace: "/v3/{p1}/3/{p2}.html"},
		{Method: http.MethodGet, Brace: "/v3/{p1}/4/{p2}.html"},
		{Method: http.MethodGet, Brace: "/v3/{p1}/5/{p2}.html"},
		{Method: http.MethodGet, Brace: "/v3/{p1}/6/{p2}.html"},
		{Method: http.MethodGet, Brace: "/v3/{p1}/7/{p2}.html"},
		{Method: http.MethodGet, Brace: "/v3/{p1}/8/{p2}.html"},
		{Method: http.MethodGet, Brace: "/v3/{p1}/9/{p2}.html"},

		{Method: http.MethodGet, Brace: "/v4/{p1}/0/{p2}.html"},
		{Method: http.MethodGet, Brace: "/v4/{p1}/1/{p2}.html"},
		{Method: http.MethodGet, Brace: "/v4/{p1}/2/{p2}.html"},
		{Method: http.MethodGet, Brace: "/v4/{p1}/3/{p2}.html"},
		{Method: http.MethodGet, Brace: "/v4/{p1}/4/{p2}.html"},
		{Method: http.MethodGet, Brace: "/v4/{p1}/5/{p2}.html"},
		{Method: http.MethodGet, Brace: "/v4/{p1}/6/{p2}.html"},
		{Method: http.MethodGet, Brace: "/v4/{p1}/7/{p2}.html"},
		{Method: http.MethodGet, Brace: "/v4/{p1}/8/{p2}.html"},
		{Method: http.MethodGet, Brace: "/v4/{p1}/9/{p2}.html"},

		// POST

		{Method: http.MethodPost, Brace: "/v0/{p1}/0/{p2}.html"},
		{Method: http.MethodPost, Brace: "/v0/{p1}/1/{p2}.html"},
		{Method: http.MethodPost, Brace: "/v0/{p1}/2/{p2}.html"},
		{Method: http.MethodPost, Brace: "/v0/{p1}/3/{p2}.html"},
		{Method: http.MethodPost, Brace: "/v0/{p1}/4/{p2}.html"},
		{Method: http.MethodPost, Brace: "/v0/{p1}/5/{p2}.html"},
		{Method: http.MethodPost, Brace: "/v0/{p1}/6/{p2}.html"},
		{Method: http.MethodPost, Brace: "/v0/{p1}/7/{p2}.html"},
		{Method: http.MethodPost, Brace: "/v0/{p1}/8/{p2}.html"},
		{Method: http.MethodPost, Brace: "/v0/{p1}/9/{p2}.html"},

		{Method: http.MethodPost, Brace: "/v1/{p1}/0/{p2}.html"},
		{Method: http.MethodPost, Brace: "/v1/{p1}/1/{p2}.html"},
		{Method: http.MethodPost, Brace: "/v1/{p1}/2/{p2}.html"},
		{Method: http.MethodPost, Brace: "/v1/{p1}/3/{p2}.html"},
		{Method: http.MethodPost, Brace: "/v1/{p1}/4/{p2}.html"},
		{Method: http.MethodPost, Brace: "/v1/{p1}/5/{p2}.html"},
		{Method: http.MethodPost, Brace: "/v1/{p1}/6/{p2}.html"},
		{Method: http.MethodPost, Brace: "/v1/{p1}/7/{p2}.html"},
		{Method: http.MethodPost, Brace: "/v1/{p1}/8/{p2}.html"},
		{Method: http.MethodPost, Brace: "/v1/{p1}/9/{p2}.html"},

		{Method: http.MethodPost, Brace: "/v2/{p1}/0/{p2}.html"},
		{Method: http.MethodPost, Brace: "/v2/{p1}/1/{p2}.html"},
		{Method: http.MethodPost, Brace: "/v2/{p1}/2/{p2}.html"},
		{Method: http.MethodPost, Brace: "/v2/{p1}/3/{p2}.html"},
		{Method: http.MethodPost, Brace: "/v2/{p1}/4/{p2}.html"},
		{Method: http.MethodPost, Brace: "/v2/{p1}/5/{p2}.html"},
		{Method: http.MethodPost, Brace: "/v2/{p1}/6/{p2}.html"},
		{Method: http.MethodPost, Brace: "/v2/{p1}/7/{p2}.html"},
		{Method: http.MethodPost, Brace: "/v2/{p1}/8/{p2}.html"},
		{Method: http.MethodPost, Brace: "/v2/{p1}/9/{p2}.html"},

		{Method: http.MethodPost, Brace: "/v3/{p1}/0/{p2}.html"},
		{Method: http.MethodPost, Brace: "/v3/{p1}/1/{p2}.html"},
		{Method: http.MethodPost, Brace: "/v3/{p1}/2/{p2}.html"},
		{Method: http.MethodPost, Brace: "/v3/{p1}/3/{p2}.html"},
		{Method: http.MethodPost, Brace: "/v3/{p1}/4/{p2}.html"},
		{Method: http.MethodPost, Brace: "/v3/{p1}/5/{p2}.html"},
		{Method: http.MethodPost, Brace: "/v3/{p1}/6/{p2}.html"},
		{Method: http.MethodPost, Brace: "/v3/{p1}/7/{p2}.html"},
		{Method: http.MethodPost, Brace: "/v3/{p1}/8/{p2}.html"},
		{Method: http.MethodPost, Brace: "/v3/{p1}/9/{p2}.html"},

		{Method: http.MethodPost, Brace: "/v4/{p1}/0/{p2}.html"},
		{Method: http.MethodPost, Brace: "/v4/{p1}/1/{p2}.html"},
		{Method: http.MethodPost, Brace: "/v4/{p1}/2/{p2}.html"},
		{Method: http.MethodPost, Brace: "/v4/{p1}/3/{p2}.html"},
		{Method: http.MethodPost, Brace: "/v4/{p1}/4/{p2}.html"},
		{Method: http.MethodPost, Brace: "/v4/{p1}/5/{p2}.html"},
		{Method: http.MethodPost, Brace: "/v4/{p1}/6/{p2}.html"},
		{Method: http.MethodPost, Brace: "/v4/{p1}/7/{p2}.html"},
		{Method: http.MethodPost, Brace: "/v4/{p1}/8/{p2}.html"},
		{Method: http.MethodPost, Brace: "/v4/{p1}/9/{p2}.html"},

		// PATCH

		{Method: http.MethodPatch, Brace: "/v0/{p1}/0/{p2}.html"},
		{Method: http.MethodPatch, Brace: "/v0/{p1}/1/{p2}.html"},
		{Method: http.MethodPatch, Brace: "/v0/{p1}/2/{p2}.html"},
		{Method: http.MethodPatch, Brace: "/v0/{p1}/3/{p2}.html"},
		{Method: http.MethodPatch, Brace: "/v0/{p1}/4/{p2}.html"},
		{Method: http.MethodPatch, Brace: "/v0/{p1}/5/{p2}.html"},
		{Method: http.MethodPatch, Brace: "/v0/{p1}/6/{p2}.html"},
		{Method: http.MethodPatch, Brace: "/v0/{p1}/7/{p2}.html"},
		{Method: http.MethodPatch, Brace: "/v0/{p1}/8/{p2}.html"},
		{Method: http.MethodPatch, Brace: "/v0/{p1}/9/{p2}.html"},

		{Method: http.MethodPatch, Brace: "/v1/{p1}/0/{p2}.html"},
		{Method: http.MethodPatch, Brace: "/v1/{p1}/1/{p2}.html"},
		{Method: http.MethodPatch, Brace: "/v1/{p1}/2/{p2}.html"},
		{Method: http.MethodPatch, Brace: "/v1/{p1}/3/{p2}.html"},
		{Method: http.MethodPatch, Brace: "/v1/{p1}/4/{p2}.html"},
		{Method: http.MethodPatch, Brace: "/v1/{p1}/5/{p2}.html"},
		{Method: http.MethodPatch, Brace: "/v1/{p1}/6/{p2}.html"},
		{Method: http.MethodPatch, Brace: "/v1/{p1}/7/{p2}.html"},
		{Method: http.MethodPatch, Brace: "/v1/{p1}/8/{p2}.html"},
		{Method: http.MethodPatch, Brace: "/v1/{p1}/9/{p2}.html"},

		{Method: http.MethodPatch, Brace: "/v2/{p1}/0/{p2}.html"},
		{Method: http.MethodPatch, Brace: "/v2/{p1}/1/{p2}.html"},
		{Method: http.MethodPatch, Brace: "/v2/{p1}/2/{p2}.html"},
		{Method: http.MethodPatch, Brace: "/v2/{p1}/3/{p2}.html"},
		{Method: http.MethodPatch, Brace: "/v2/{p1}/4/{p2}.html"},
		{Method: http.MethodPatch, Brace: "/v2/{p1}/5/{p2}.html"},
		{Method: http.MethodPatch, Brace: "/v2/{p1}/6/{p2}.html"},
		{Method: http.MethodPatch, Brace: "/v2/{p1}/7/{p2}.html"},
		{Method: http.MethodPatch, Brace: "/v2/{p1}/8/{p2}.html"},
		{Method: http.MethodPatch, Brace: "/v2/{p1}/9/{p2}.html"},

		{Method: http.MethodPatch, Brace: "/v3/{p1}/0/{p2}.html"},
		{Method: http.MethodPatch, Brace: "/v3/{p1}/1/{p2}.html"},
		{Method: http.MethodPatch, Brace: "/v3/{p1}/2/{p2}.html"},
		{Method: http.MethodPatch, Brace: "/v3/{p1}/3/{p2}.html"},
		{Method: http.MethodPatch, Brace: "/v3/{p1}/4/{p2}.html"},
		{Method: http.MethodPatch, Brace: "/v3/{p1}/5/{p2}.html"},
		{Method: http.MethodPatch, Brace: "/v3/{p1}/6/{p2}.html"},
		{Method: http.MethodPatch, Brace: "/v3/{p1}/7/{p2}.html"},
		{Method: http.MethodPatch, Brace: "/v3/{p1}/8/{p2}.html"},
		{Method: http.MethodPatch, Brace: "/v3/{p1}/9/{p2}.html"},

		{Method: http.MethodPatch, Brace: "/v4/{p1}/0/{p2}.html"},
		{Method: http.MethodPatch, Brace: "/v4/{p1}/1/{p2}.html"},
		{Method: http.MethodPatch, Brace: "/v4/{p1}/2/{p2}.html"},
		{Method: http.MethodPatch, Brace: "/v4/{p1}/3/{p2}.html"},
		{Method: http.MethodPatch, Brace: "/v4/{p1}/4/{p2}.html"},
		{Method: http.MethodPatch, Brace: "/v4/{p1}/5/{p2}.html"},
		{Method: http.MethodPatch, Brace: "/v4/{p1}/6/{p2}.html"},
		{Method: http.MethodPatch, Brace: "/v4/{p1}/7/{p2}.html"},
		{Method: http.MethodPatch, Brace: "/v4/{p1}/8/{p2}.html"},
		{Method: http.MethodPatch, Brace: "/v4/{p1}/9/{p2}.html"},
	},
}
