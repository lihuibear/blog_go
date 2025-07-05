package cmd

import (
	"blog/internal/controller/admin"
	"blog/internal/controller/article"
	"blog/internal/controller/article_grp"
	"context"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gcmd"

	"blog/internal/controller/hello"
)

func MiddlewareCORS(r *ghttp.Request) {
	r.Response.CORSDefault()
	r.Middleware.Next()
}

var (
	Main = gcmd.Command{
		Name:  "main",
		Usage: "main",
		Brief: "start http server",
		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
			s := g.Server()

			s.Group("/", func(group *ghttp.RouterGroup) {
				group.Middleware(ghttp.MiddlewareHandlerResponse)
				group.Group("/v1", func(group *ghttp.RouterGroup) {
					group.Middleware(MiddlewareCORS)
					group.Bind(
						hello.NewV1(),
						admin.NewV1(),
						article_grp.NewV1(),
						article.NewV1(),
					)
				})

			})
			s.Run()
			s.SetPort(8080)
			return nil
		},
	}
)
