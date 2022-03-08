package router

import (
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
)

func init() {
	s := g.Server()
	s.BindHandler("/test", func(req *ghttp.Request) {
		req.Response.Write("test")
	})
}
