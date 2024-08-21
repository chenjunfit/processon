package router

import (
	"context"
	"github.com/gogf/gf/v2/net/ghttp"
	"processon/internal/app/server/controller"
)

var R = new(Router)

type Router struct{}

func (router *Router) BindController(ctx context.Context, group *ghttp.RouterGroup) {
	group.Group("/processon", func(group *ghttp.RouterGroup) {
		group.Middleware(func(r *ghttp.Request) {
			r.SetCtx(r.GetNeverDoneCtx())
			r.Middleware.Next()
		})
		group.Bind(
			controller.BaseLine,
			controller.CheckScript,
			controller.CheckJob,
			controller.FailedResult,
		)
	})

}
