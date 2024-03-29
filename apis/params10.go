// SPDX-FileCopyrightText: 2017-2024 caixw
//
// SPDX-License-Identifier: MIT

package apis

import "net/http"

func init() {
	params10.init()
}

var params10 = &Collection{
	ID:   "10-params",
	Name: "10 Params",
	Desc: "十个参数的路由",
	APIS: []*API{
		{Method: http.MethodGet, Brace: "/v0/{p1}/00/{p2}/00/{p3}/0000/{p4}/00000/{p5}/1/{p6}/1/{p7}/1/{p8}/1/{p9}/{p10}.html"},
		{Method: http.MethodGet, Brace: "/v0/{p1}/01/{p2}/00/{p3}/0000/{p4}/00000/{p5}/1/{p6}/1/{p7}/1/{p8}/1/{p9}/{p10}.html"},
		{Method: http.MethodGet, Brace: "/v0/{p1}/02/{p2}/00/{p3}/0000/{p4}/00000/{p5}/1/{p6}/1/{p7}/1/{p8}/1/{p9}/{p10}.html"},
		{Method: http.MethodGet, Brace: "/v0/{p1}/03/{p2}/00/{p3}/0000/{p4}/00000/{p5}/1/{p6}/1/{p7}/1/{p8}/1/{p9}/{p10}.html"},
		{Method: http.MethodGet, Brace: "/v0/{p1}/04/{p2}/00/{p3}/0000/{p4}/00000/{p5}/1/{p6}/1/{p7}/1/{p8}/1/{p9}/{p10}.html"},
		{Method: http.MethodGet, Brace: "/v0/{p1}/05/{p2}/00/{p3}/0000/{p4}/00000/{p5}/1/{p6}/1/{p7}/1/{p8}/1/{p9}/{p10}.html"},
		{Method: http.MethodGet, Brace: "/v0/{p1}/06/{p2}/00/{p3}/0000/{p4}/00000/{p5}/1/{p6}/1/{p7}/1/{p8}/1/{p9}/{p10}.html"},
		{Method: http.MethodGet, Brace: "/v0/{p1}/07/{p2}/00/{p3}/0000/{p4}/00000/{p5}/1/{p6}/1/{p7}/1/{p8}/1/{p9}/{p10}.html"},
		{Method: http.MethodGet, Brace: "/v0/{p1}/08/{p2}/00/{p3}/0000/{p4}/00000/{p5}/1/{p6}/1/{p7}/1/{p8}/1/{p9}/{p10}.html"},
		{Method: http.MethodGet, Brace: "/v0/{p1}/09/{p2}/00/{p3}/0000/{p4}/00000/{p5}/1/{p6}/1/{p7}/1/{p8}/1/{p9}/{p10}.html"},

		{Method: http.MethodGet, Brace: "/v0/{p1}/10/{p2}/00/{p3}/0000/{p4}/00000/{p5}/1/{p6}/1/{p7}/1/{p8}/1/{p9}/{p10}.html"},
		{Method: http.MethodGet, Brace: "/v0/{p1}/11/{p2}/00/{p3}/0000/{p4}/00000/{p5}/1/{p6}/1/{p7}/1/{p8}/1/{p9}/{p10}.html"},
		{Method: http.MethodGet, Brace: "/v0/{p1}/12/{p2}/00/{p3}/0000/{p4}/00000/{p5}/1/{p6}/1/{p7}/1/{p8}/1/{p9}/{p10}.html"},
		{Method: http.MethodGet, Brace: "/v0/{p1}/13/{p2}/00/{p3}/0000/{p4}/00000/{p5}/1/{p6}/1/{p7}/1/{p8}/1/{p9}/{p10}.html"},
		{Method: http.MethodGet, Brace: "/v0/{p1}/14/{p2}/00/{p3}/0000/{p4}/00000/{p5}/1/{p6}/1/{p7}/1/{p8}/1/{p9}/{p10}.html"},
		{Method: http.MethodGet, Brace: "/v0/{p1}/15/{p2}/00/{p3}/0000/{p4}/00000/{p5}/1/{p6}/1/{p7}/1/{p8}/1/{p9}/{p10}.html"},
		{Method: http.MethodGet, Brace: "/v0/{p1}/16/{p2}/00/{p3}/0000/{p4}/00000/{p5}/1/{p6}/1/{p7}/1/{p8}/1/{p9}/{p10}.html"},
		{Method: http.MethodGet, Brace: "/v0/{p1}/17/{p2}/00/{p3}/0000/{p4}/00000/{p5}/1/{p6}/1/{p7}/1/{p8}/1/{p9}/{p10}.html"},
		{Method: http.MethodGet, Brace: "/v0/{p1}/18/{p2}/00/{p3}/0000/{p4}/00000/{p5}/1/{p6}/1/{p7}/1/{p8}/1/{p9}/{p10}.html"},
		{Method: http.MethodGet, Brace: "/v0/{p1}/19/{p2}/00/{p3}/0000/{p4}/00000/{p5}/1/{p6}/1/{p7}/1/{p8}/1/{p9}/{p10}.html"},

		{Method: http.MethodGet, Brace: "/v0/{p1}/10/{p2}/10/{p3}/0000/{p4}/00000/{p5}/1/{p6}/1/{p7}/1/{p8}/1/{p9}/{p10}.html"},
		{Method: http.MethodGet, Brace: "/v0/{p1}/11/{p2}/01/{p3}/0000/{p4}/00000/{p5}/1/{p6}/1/{p7}/1/{p8}/1/{p9}/{p10}.html"},
		{Method: http.MethodGet, Brace: "/v0/{p1}/12/{p2}/02/{p3}/0000/{p4}/00000/{p5}/1/{p6}/1/{p7}/1/{p8}/1/{p9}/{p10}.html"},
		{Method: http.MethodGet, Brace: "/v0/{p1}/13/{p2}/03/{p3}/0000/{p4}/00000/{p5}/1/{p6}/1/{p7}/1/{p8}/1/{p9}/{p10}.html"},
		{Method: http.MethodGet, Brace: "/v0/{p1}/14/{p2}/04/{p3}/0000/{p4}/00000/{p5}/1/{p6}/1/{p7}/1/{p8}/1/{p9}/{p10}.html"},
		{Method: http.MethodGet, Brace: "/v0/{p1}/15/{p2}/05/{p3}/0000/{p4}/00000/{p5}/1/{p6}/1/{p7}/1/{p8}/1/{p9}/{p10}.html"},
		{Method: http.MethodGet, Brace: "/v0/{p1}/16/{p2}/06/{p3}/0000/{p4}/00000/{p5}/1/{p6}/1/{p7}/1/{p8}/1/{p9}/{p10}.html"},
		{Method: http.MethodGet, Brace: "/v0/{p1}/17/{p2}/07/{p3}/0000/{p4}/00000/{p5}/1/{p6}/1/{p7}/1/{p8}/1/{p9}/{p10}.html"},
		{Method: http.MethodGet, Brace: "/v0/{p1}/18/{p2}/08/{p3}/0000/{p4}/00000/{p5}/1/{p6}/1/{p7}/1/{p8}/1/{p9}/{p10}.html"},
		{Method: http.MethodGet, Brace: "/v0/{p1}/19/{p2}/09/{p3}/0000/{p4}/00000/{p5}/1/{p6}/1/{p7}/1/{p8}/1/{p9}/{p10}.html"},

		{Method: http.MethodGet, Brace: "/v0/{p1}/20/{p2}/00/{p3}/0000/{p4}/00000/{p5}/1/{p6}/1/{p7}/1/{p8}/1/{p9}/{p10}.html"},
		{Method: http.MethodGet, Brace: "/v0/{p1}/21/{p2}/01/{p3}/0000/{p4}/00000/{p5}/1/{p6}/1/{p7}/1/{p8}/1/{p9}/{p10}.html"},
		{Method: http.MethodGet, Brace: "/v0/{p1}/22/{p2}/02/{p3}/0000/{p4}/00000/{p5}/1/{p6}/1/{p7}/1/{p8}/1/{p9}/{p10}.html"},
		{Method: http.MethodGet, Brace: "/v0/{p1}/23/{p2}/03/{p3}/0000/{p4}/00000/{p5}/1/{p6}/1/{p7}/1/{p8}/1/{p9}/{p10}.html"},
		{Method: http.MethodGet, Brace: "/v0/{p1}/24/{p2}/04/{p3}/0000/{p4}/00000/{p5}/1/{p6}/1/{p7}/1/{p8}/1/{p9}/{p10}.html"},
		{Method: http.MethodGet, Brace: "/v0/{p1}/25/{p2}/05/{p3}/0000/{p4}/00000/{p5}/1/{p6}/1/{p7}/1/{p8}/1/{p9}/{p10}.html"},
		{Method: http.MethodGet, Brace: "/v0/{p1}/26/{p2}/06/{p3}/0000/{p4}/00000/{p5}/1/{p6}/1/{p7}/1/{p8}/1/{p9}/{p10}.html"},
		{Method: http.MethodGet, Brace: "/v0/{p1}/27/{p2}/07/{p3}/0000/{p4}/00000/{p5}/1/{p6}/1/{p7}/1/{p8}/1/{p9}/{p10}.html"},
		{Method: http.MethodGet, Brace: "/v0/{p1}/28/{p2}/08/{p3}/0000/{p4}/00000/{p5}/1/{p6}/1/{p7}/1/{p8}/1/{p9}/{p10}.html"},
		{Method: http.MethodGet, Brace: "/v0/{p1}/29/{p2}/09/{p3}/0000/{p4}/00000/{p5}/1/{p6}/1/{p7}/1/{p8}/1/{p9}/{p10}.html"},

		// PUT

		{Method: http.MethodPut, Brace: "/v0/{p1}/00/{p2}/00/{p3}/0000/{p4}/00000/{p5}/1/{p6}/1/{p7}/1/{p8}/1/{p9}/{p10}.html"},
		{Method: http.MethodPut, Brace: "/v0/{p1}/01/{p2}/00/{p3}/0000/{p4}/00000/{p5}/1/{p6}/1/{p7}/1/{p8}/1/{p9}/{p10}.html"},
		{Method: http.MethodPut, Brace: "/v0/{p1}/02/{p2}/00/{p3}/0000/{p4}/00000/{p5}/1/{p6}/1/{p7}/1/{p8}/1/{p9}/{p10}.html"},
		{Method: http.MethodPut, Brace: "/v0/{p1}/03/{p2}/00/{p3}/0000/{p4}/00000/{p5}/1/{p6}/1/{p7}/1/{p8}/1/{p9}/{p10}.html"},
		{Method: http.MethodPut, Brace: "/v0/{p1}/04/{p2}/00/{p3}/0000/{p4}/00000/{p5}/1/{p6}/1/{p7}/1/{p8}/1/{p9}/{p10}.html"},
		{Method: http.MethodPut, Brace: "/v0/{p1}/05/{p2}/00/{p3}/0000/{p4}/00000/{p5}/1/{p6}/1/{p7}/1/{p8}/1/{p9}/{p10}.html"},
		{Method: http.MethodPut, Brace: "/v0/{p1}/06/{p2}/00/{p3}/0000/{p4}/00000/{p5}/1/{p6}/1/{p7}/1/{p8}/1/{p9}/{p10}.html"},
		{Method: http.MethodPut, Brace: "/v0/{p1}/07/{p2}/00/{p3}/0000/{p4}/00000/{p5}/1/{p6}/1/{p7}/1/{p8}/1/{p9}/{p10}.html"},
		{Method: http.MethodPut, Brace: "/v0/{p1}/08/{p2}/00/{p3}/0000/{p4}/00000/{p5}/1/{p6}/1/{p7}/1/{p8}/1/{p9}/{p10}.html"},
		{Method: http.MethodPut, Brace: "/v0/{p1}/09/{p2}/00/{p3}/0000/{p4}/00000/{p5}/1/{p6}/1/{p7}/1/{p8}/1/{p9}/{p10}.html"},

		{Method: http.MethodPut, Brace: "/v0/{p1}/10/{p2}/00/{p3}/0000/{p4}/00000/{p5}/1/{p6}/1/{p7}/1/{p8}/1/{p9}/{p10}.html"},
		{Method: http.MethodPut, Brace: "/v0/{p1}/11/{p2}/00/{p3}/0000/{p4}/00000/{p5}/1/{p6}/1/{p7}/1/{p8}/1/{p9}/{p10}.html"},
		{Method: http.MethodPut, Brace: "/v0/{p1}/12/{p2}/00/{p3}/0000/{p4}/00000/{p5}/1/{p6}/1/{p7}/1/{p8}/1/{p9}/{p10}.html"},
		{Method: http.MethodPut, Brace: "/v0/{p1}/13/{p2}/00/{p3}/0000/{p4}/00000/{p5}/1/{p6}/1/{p7}/1/{p8}/1/{p9}/{p10}.html"},
		{Method: http.MethodPut, Brace: "/v0/{p1}/14/{p2}/00/{p3}/0000/{p4}/00000/{p5}/1/{p6}/1/{p7}/1/{p8}/1/{p9}/{p10}.html"},
		{Method: http.MethodPut, Brace: "/v0/{p1}/15/{p2}/00/{p3}/0000/{p4}/00000/{p5}/1/{p6}/1/{p7}/1/{p8}/1/{p9}/{p10}.html"},
		{Method: http.MethodPut, Brace: "/v0/{p1}/16/{p2}/00/{p3}/0000/{p4}/00000/{p5}/1/{p6}/1/{p7}/1/{p8}/1/{p9}/{p10}.html"},
		{Method: http.MethodPut, Brace: "/v0/{p1}/17/{p2}/00/{p3}/0000/{p4}/00000/{p5}/1/{p6}/1/{p7}/1/{p8}/1/{p9}/{p10}.html"},
		{Method: http.MethodPut, Brace: "/v0/{p1}/18/{p2}/00/{p3}/0000/{p4}/00000/{p5}/1/{p6}/1/{p7}/1/{p8}/1/{p9}/{p10}.html"},
		{Method: http.MethodPut, Brace: "/v0/{p1}/19/{p2}/00/{p3}/0000/{p4}/00000/{p5}/1/{p6}/1/{p7}/1/{p8}/1/{p9}/{p10}.html"},

		{Method: http.MethodPut, Brace: "/v0/{p1}/10/{p2}/10/{p3}/0000/{p4}/00000/{p5}/1/{p6}/1/{p7}/1/{p8}/1/{p9}/{p10}.html"},
		{Method: http.MethodPut, Brace: "/v0/{p1}/11/{p2}/01/{p3}/0000/{p4}/00000/{p5}/1/{p6}/1/{p7}/1/{p8}/1/{p9}/{p10}.html"},
		{Method: http.MethodPut, Brace: "/v0/{p1}/12/{p2}/02/{p3}/0000/{p4}/00000/{p5}/1/{p6}/1/{p7}/1/{p8}/1/{p9}/{p10}.html"},
		{Method: http.MethodPut, Brace: "/v0/{p1}/13/{p2}/03/{p3}/0000/{p4}/00000/{p5}/1/{p6}/1/{p7}/1/{p8}/1/{p9}/{p10}.html"},
		{Method: http.MethodPut, Brace: "/v0/{p1}/14/{p2}/04/{p3}/0000/{p4}/00000/{p5}/1/{p6}/1/{p7}/1/{p8}/1/{p9}/{p10}.html"},
		{Method: http.MethodPut, Brace: "/v0/{p1}/15/{p2}/05/{p3}/0000/{p4}/00000/{p5}/1/{p6}/1/{p7}/1/{p8}/1/{p9}/{p10}.html"},
		{Method: http.MethodPut, Brace: "/v0/{p1}/16/{p2}/06/{p3}/0000/{p4}/00000/{p5}/1/{p6}/1/{p7}/1/{p8}/1/{p9}/{p10}.html"},
		{Method: http.MethodPut, Brace: "/v0/{p1}/17/{p2}/07/{p3}/0000/{p4}/00000/{p5}/1/{p6}/1/{p7}/1/{p8}/1/{p9}/{p10}.html"},
		{Method: http.MethodPut, Brace: "/v0/{p1}/18/{p2}/08/{p3}/0000/{p4}/00000/{p5}/1/{p6}/1/{p7}/1/{p8}/1/{p9}/{p10}.html"},
		{Method: http.MethodPut, Brace: "/v0/{p1}/19/{p2}/09/{p3}/0000/{p4}/00000/{p5}/1/{p6}/1/{p7}/1/{p8}/1/{p9}/{p10}.html"},

		{Method: http.MethodPut, Brace: "/v0/{p1}/20/{p2}/00/{p3}/0000/{p4}/00000/{p5}/1/{p6}/1/{p7}/1/{p8}/1/{p9}/{p10}.html"},
		{Method: http.MethodPut, Brace: "/v0/{p1}/21/{p2}/01/{p3}/0000/{p4}/00000/{p5}/1/{p6}/1/{p7}/1/{p8}/1/{p9}/{p10}.html"},
		{Method: http.MethodPut, Brace: "/v0/{p1}/22/{p2}/02/{p3}/0000/{p4}/00000/{p5}/1/{p6}/1/{p7}/1/{p8}/1/{p9}/{p10}.html"},
		{Method: http.MethodPut, Brace: "/v0/{p1}/23/{p2}/03/{p3}/0000/{p4}/00000/{p5}/1/{p6}/1/{p7}/1/{p8}/1/{p9}/{p10}.html"},
		{Method: http.MethodPut, Brace: "/v0/{p1}/24/{p2}/04/{p3}/0000/{p4}/00000/{p5}/1/{p6}/1/{p7}/1/{p8}/1/{p9}/{p10}.html"},
		{Method: http.MethodPut, Brace: "/v0/{p1}/25/{p2}/05/{p3}/0000/{p4}/00000/{p5}/1/{p6}/1/{p7}/1/{p8}/1/{p9}/{p10}.html"},
		{Method: http.MethodPut, Brace: "/v0/{p1}/26/{p2}/06/{p3}/0000/{p4}/00000/{p5}/1/{p6}/1/{p7}/1/{p8}/1/{p9}/{p10}.html"},
		{Method: http.MethodPut, Brace: "/v0/{p1}/27/{p2}/07/{p3}/0000/{p4}/00000/{p5}/1/{p6}/1/{p7}/1/{p8}/1/{p9}/{p10}.html"},
		{Method: http.MethodPut, Brace: "/v0/{p1}/28/{p2}/08/{p3}/0000/{p4}/00000/{p5}/1/{p6}/1/{p7}/1/{p8}/1/{p9}/{p10}.html"},
		{Method: http.MethodPut, Brace: "/v0/{p1}/29/{p2}/09/{p3}/0000/{p4}/00000/{p5}/1/{p6}/1/{p7}/1/{p8}/1/{p9}/{p10}.html"},

		// POST

		{Method: http.MethodPost, Brace: "/v0/{p1}/00/{p2}/00/{p3}/0000/{p4}/00000/{p5}/1/{p6}/1/{p7}/1/{p8}/1/{p9}/{p10}.html"},
		{Method: http.MethodPost, Brace: "/v0/{p1}/01/{p2}/00/{p3}/0000/{p4}/00000/{p5}/1/{p6}/1/{p7}/1/{p8}/1/{p9}/{p10}.html"},
		{Method: http.MethodPost, Brace: "/v0/{p1}/02/{p2}/00/{p3}/0000/{p4}/00000/{p5}/1/{p6}/1/{p7}/1/{p8}/1/{p9}/{p10}.html"},
		{Method: http.MethodPost, Brace: "/v0/{p1}/03/{p2}/00/{p3}/0000/{p4}/00000/{p5}/1/{p6}/1/{p7}/1/{p8}/1/{p9}/{p10}.html"},
		{Method: http.MethodPost, Brace: "/v0/{p1}/04/{p2}/00/{p3}/0000/{p4}/00000/{p5}/1/{p6}/1/{p7}/1/{p8}/1/{p9}/{p10}.html"},
		{Method: http.MethodPost, Brace: "/v0/{p1}/05/{p2}/00/{p3}/0000/{p4}/00000/{p5}/1/{p6}/1/{p7}/1/{p8}/1/{p9}/{p10}.html"},
		{Method: http.MethodPost, Brace: "/v0/{p1}/06/{p2}/00/{p3}/0000/{p4}/00000/{p5}/1/{p6}/1/{p7}/1/{p8}/1/{p9}/{p10}.html"},
		{Method: http.MethodPost, Brace: "/v0/{p1}/07/{p2}/00/{p3}/0000/{p4}/00000/{p5}/1/{p6}/1/{p7}/1/{p8}/1/{p9}/{p10}.html"},
		{Method: http.MethodPost, Brace: "/v0/{p1}/08/{p2}/00/{p3}/0000/{p4}/00000/{p5}/1/{p6}/1/{p7}/1/{p8}/1/{p9}/{p10}.html"},
		{Method: http.MethodPost, Brace: "/v0/{p1}/09/{p2}/00/{p3}/0000/{p4}/00000/{p5}/1/{p6}/1/{p7}/1/{p8}/1/{p9}/{p10}.html"},

		{Method: http.MethodPost, Brace: "/v0/{p1}/10/{p2}/00/{p3}/0000/{p4}/00000/{p5}/1/{p6}/1/{p7}/1/{p8}/1/{p9}/{p10}.html"},
		{Method: http.MethodPost, Brace: "/v0/{p1}/11/{p2}/00/{p3}/0000/{p4}/00000/{p5}/1/{p6}/1/{p7}/1/{p8}/1/{p9}/{p10}.html"},
		{Method: http.MethodPost, Brace: "/v0/{p1}/12/{p2}/00/{p3}/0000/{p4}/00000/{p5}/1/{p6}/1/{p7}/1/{p8}/1/{p9}/{p10}.html"},
		{Method: http.MethodPost, Brace: "/v0/{p1}/13/{p2}/00/{p3}/0000/{p4}/00000/{p5}/1/{p6}/1/{p7}/1/{p8}/1/{p9}/{p10}.html"},
		{Method: http.MethodPost, Brace: "/v0/{p1}/14/{p2}/00/{p3}/0000/{p4}/00000/{p5}/1/{p6}/1/{p7}/1/{p8}/1/{p9}/{p10}.html"},
		{Method: http.MethodPost, Brace: "/v0/{p1}/15/{p2}/00/{p3}/0000/{p4}/00000/{p5}/1/{p6}/1/{p7}/1/{p8}/1/{p9}/{p10}.html"},
		{Method: http.MethodPost, Brace: "/v0/{p1}/16/{p2}/00/{p3}/0000/{p4}/00000/{p5}/1/{p6}/1/{p7}/1/{p8}/1/{p9}/{p10}.html"},
		{Method: http.MethodPost, Brace: "/v0/{p1}/17/{p2}/00/{p3}/0000/{p4}/00000/{p5}/1/{p6}/1/{p7}/1/{p8}/1/{p9}/{p10}.html"},
		{Method: http.MethodPost, Brace: "/v0/{p1}/18/{p2}/00/{p3}/0000/{p4}/00000/{p5}/1/{p6}/1/{p7}/1/{p8}/1/{p9}/{p10}.html"},
		{Method: http.MethodPost, Brace: "/v0/{p1}/19/{p2}/00/{p3}/0000/{p4}/00000/{p5}/1/{p6}/1/{p7}/1/{p8}/1/{p9}/{p10}.html"},

		{Method: http.MethodPost, Brace: "/v0/{p1}/10/{p2}/10/{p3}/0000/{p4}/00000/{p5}/1/{p6}/1/{p7}/1/{p8}/1/{p9}/{p10}.html"},
		{Method: http.MethodPost, Brace: "/v0/{p1}/11/{p2}/01/{p3}/0000/{p4}/00000/{p5}/1/{p6}/1/{p7}/1/{p8}/1/{p9}/{p10}.html"},
		{Method: http.MethodPost, Brace: "/v0/{p1}/12/{p2}/02/{p3}/0000/{p4}/00000/{p5}/1/{p6}/1/{p7}/1/{p8}/1/{p9}/{p10}.html"},
		{Method: http.MethodPost, Brace: "/v0/{p1}/13/{p2}/03/{p3}/0000/{p4}/00000/{p5}/1/{p6}/1/{p7}/1/{p8}/1/{p9}/{p10}.html"},
		{Method: http.MethodPost, Brace: "/v0/{p1}/14/{p2}/04/{p3}/0000/{p4}/00000/{p5}/1/{p6}/1/{p7}/1/{p8}/1/{p9}/{p10}.html"},
		{Method: http.MethodPost, Brace: "/v0/{p1}/15/{p2}/05/{p3}/0000/{p4}/00000/{p5}/1/{p6}/1/{p7}/1/{p8}/1/{p9}/{p10}.html"},
		{Method: http.MethodPost, Brace: "/v0/{p1}/16/{p2}/06/{p3}/0000/{p4}/00000/{p5}/1/{p6}/1/{p7}/1/{p8}/1/{p9}/{p10}.html"},
		{Method: http.MethodPost, Brace: "/v0/{p1}/17/{p2}/07/{p3}/0000/{p4}/00000/{p5}/1/{p6}/1/{p7}/1/{p8}/1/{p9}/{p10}.html"},
		{Method: http.MethodPost, Brace: "/v0/{p1}/18/{p2}/08/{p3}/0000/{p4}/00000/{p5}/1/{p6}/1/{p7}/1/{p8}/1/{p9}/{p10}.html"},
		{Method: http.MethodPost, Brace: "/v0/{p1}/19/{p2}/09/{p3}/0000/{p4}/00000/{p5}/1/{p6}/1/{p7}/1/{p8}/1/{p9}/{p10}.html"},

		{Method: http.MethodPost, Brace: "/v0/{p1}/20/{p2}/00/{p3}/0000/{p4}/00000/{p5}/1/{p6}/1/{p7}/1/{p8}/1/{p9}/{p10}.html"},
		{Method: http.MethodPost, Brace: "/v0/{p1}/21/{p2}/01/{p3}/0000/{p4}/00000/{p5}/1/{p6}/1/{p7}/1/{p8}/1/{p9}/{p10}.html"},
		{Method: http.MethodPost, Brace: "/v0/{p1}/22/{p2}/02/{p3}/0000/{p4}/00000/{p5}/1/{p6}/1/{p7}/1/{p8}/1/{p9}/{p10}.html"},
		{Method: http.MethodPost, Brace: "/v0/{p1}/23/{p2}/03/{p3}/0000/{p4}/00000/{p5}/1/{p6}/1/{p7}/1/{p8}/1/{p9}/{p10}.html"},
		{Method: http.MethodPost, Brace: "/v0/{p1}/24/{p2}/04/{p3}/0000/{p4}/00000/{p5}/1/{p6}/1/{p7}/1/{p8}/1/{p9}/{p10}.html"},
		{Method: http.MethodPost, Brace: "/v0/{p1}/25/{p2}/05/{p3}/0000/{p4}/00000/{p5}/1/{p6}/1/{p7}/1/{p8}/1/{p9}/{p10}.html"},
		{Method: http.MethodPost, Brace: "/v0/{p1}/26/{p2}/06/{p3}/0000/{p4}/00000/{p5}/1/{p6}/1/{p7}/1/{p8}/1/{p9}/{p10}.html"},
		{Method: http.MethodPost, Brace: "/v0/{p1}/27/{p2}/07/{p3}/0000/{p4}/00000/{p5}/1/{p6}/1/{p7}/1/{p8}/1/{p9}/{p10}.html"},
		{Method: http.MethodPost, Brace: "/v0/{p1}/28/{p2}/08/{p3}/0000/{p4}/00000/{p5}/1/{p6}/1/{p7}/1/{p8}/1/{p9}/{p10}.html"},
		{Method: http.MethodPost, Brace: "/v0/{p1}/29/{p2}/09/{p3}/0000/{p4}/00000/{p5}/1/{p6}/1/{p7}/1/{p8}/1/{p9}/{p10}.html"},
	},
}
