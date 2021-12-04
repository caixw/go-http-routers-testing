// SPDX-License-Identifier: MIT

package routers

import (
	"github.com/fasthttp/router"
	"github.com/valyala/fasthttp"

	"github.com/caixw/go-http-routers-testing/apis"
)

func init() {
	Routers = append(Routers, &Router{
		Name: "fasthttp-router",
		URL:  "https://github.com/fasthttp/router",
		Load: fasthttpRouterLoad,
	})
}

func fasthttpRouterLoad(as []*apis.API) ServeFunc {
	r := router.New()
	for _, api := range as {
		r.Handle(api.Method, api.Brace, func(ctx *fasthttp.RequestCtx) {
			//_ = ctx.UserValue()
			ctx.Write(ctx.URI().Path())
		})
	}

	return func(api *apis.API) string {
		resp := fasthttp.AcquireResponse()
		req := fasthttp.AcquireRequest()
		req.Header.SetMethod(api.Method)
		req.SetRequestURI(api.Test)

		ctx := &fasthttp.RequestCtx{Request: *req, Response: *resp}
		r.Handler(ctx)
		body := ctx.Response.Body()

		fasthttp.ReleaseRequest(req)
		fasthttp.ReleaseResponse(resp)

		return string(body)
	}
}
