package router

import (
	"github.com/gogf/gf/net/ghttp"
	"project/main/module/user/controller"
	"project/main/tool"
)

func Register(server *ghttp.Server) {
	tool.BindObjectReflect("/user", controller.UserController, server)
}
