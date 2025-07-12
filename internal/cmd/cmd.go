package cmd

import (
	"blog/internal/controller/admin"
	"blog/internal/controller/article"
	"blog/internal/controller/article_grp"
	"blog/internal/controller/file"
	"blog/internal/controller/link"
	"blog/internal/controller/other"
	"blog/internal/controller/reading"
	"blog/internal/controller/reply"
	"blog/internal/controller/sentence"
	"blog/internal/controller/tag"
	"blog/internal/controller/tag_grp"
	"context"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gcmd"

	"blog/internal/controller/hello"
)

const (
	MySwaggerUITemplate = `
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="utf-8" />
    <meta name="viewport" content="width=device-width, initial-scale=1" />
    <meta name="description" content="SwaggerUI"/>
    <title>SwaggerUI</title>
    <link rel="stylesheet" href="https://cdnjs.cloudflare.com/ajax/libs/swagger-ui/5.10.5/swagger-ui.min.css" />
</head>
<body>
<div id="swagger-ui"></div>
<script src="https://cdnjs.cloudflare.com/ajax/libs/swagger-ui/5.10.5/swagger-ui-bundle.js" crossorigin></script>
<script>
    window.onload = () => {
        window.ui = SwaggerUIBundle({
            url:    '{SwaggerUIDocUrl}',
            dom_id: '#swagger-ui',
        });
    };
</script>
</body>
</html>
`

	OpenAPIUI = `<!doctype html>
<html lang="en">
  <head>
    <meta charset="UTF-8" />
    <title>openAPI UI</title>
  </head>
  <body>
    <div id="openapi-ui-container" spec-url="{SwaggerUIDocUrl}" theme="light"></div>
    <script src="https://cdn.jsdelivr.net/npm/openapi-ui-dist@latest/lib/openapi-ui.umd.js"></script>
  </body>
</html>`
)

func MiddlewareCORS(r *ghttp.Request) {
	// 允许所有来源的跨域请求
	r.Response.Header().Set("Access-Control-Allow-Origin", "*")
	// 设置允许的请求方法
	r.Response.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
	// 设置允许的请求头
	r.Response.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
	// 预检请求直接返回204
	if r.Method == "OPTIONS" {
		r.Response.WriteStatusExit(204)
	}
	r.Middleware.Next()
}

// Main 表示应用程序的入口命令。
// 它用于初始化并启动 HTTP 服务器，并在此过程中设置必要的路由和中间件。
var (
	Main = gcmd.Command{
		Name:  "main",
		Usage: "main",
		Brief: "启动 HTTP 服务",
		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
			// 初始化服务器实例
			s := g.Server()

			// 配置根路由组
			s.Group("/", func(group *ghttp.RouterGroup) {
				// 添加全局中间件：HandlerResponse，用于统一处理响应格式
				group.Middleware(ghttp.MiddlewareHandlerResponse)
				// 配置 /admin 路由组
				group.Group("/admin", func(group *ghttp.RouterGroup) {
					// 添加跨域请求处理中间件
					group.Middleware(MiddlewareCORS)
					// 绑定控制器，注册各个业务接口
					group.Bind(
						admin.NewV1(), // 管理员模块 V1 接口
					)
				})

				// 配置 /v1 路由组
				group.Group("/v1", func(group *ghttp.RouterGroup) {
					// 添加跨域请求处理中间件
					group.Middleware(MiddlewareCORS)

					// 绑定控制器，注册各个业务接口
					group.Bind(
						hello.NewV1(),       // Hello 模块 V1 接口
						article_grp.NewV1(), // 文章分组模块 V1 接口
						article.NewV1(),     // 文章模块 V1 接口
						tag_grp.NewV1(),     // 标签分组模块 V1 接口
						tag.NewV1(),         // 标签模块 V1 接口
						file.NewV1(),
						link.NewV1(),
						reply.NewV1(),
						reading.NewV1(),
						sentence.NewV1(),
					)
				})
				// 配置 /v1 路由组
				group.Group("/app", func(group *ghttp.RouterGroup) {
					// 添加跨域请求处理中间件
					group.Middleware(MiddlewareCORS)

					// 绑定控制器，注册各个业务接口
					group.Bind(
						article.NewApp(),
						article_grp.NewApp(),
						reply.NewApp(),
						other.NewApp(),
					)
				})

			})

			// 设置服务器监听端口
			s.SetPort(8000)
			s.SetSwaggerUITemplate(MySwaggerUITemplate)
			// 启动服务器
			s.Run()

			// 返回 nil 表示命令执行成功
			return nil
		},
	}
)
