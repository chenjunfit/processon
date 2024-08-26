package cmd

import (
	"context"
	_ "github.com/gogf/gf/contrib/drivers/mysql/v2"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gcmd"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"processon/internal/app/server/router"
)

var (
	Http = gcmd.Command{
		Name:  "main",
		Usage: "main",
		Brief: "start http server",
		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
			s := g.Server()
			s.SetIndexFolder(true)
			s.SetServerRoot("resource/agent/")
			s.Group("/", func(group *ghttp.RouterGroup) {
				group.Middleware(ghttp.MiddlewareHandlerResponse)
				router.R.BindController(ctx, group)

			})
			s.BindHandler("/metrics", ghttp.WrapH(promhttp.Handler()))
			s.Run()
			return nil
		},
	}
)
