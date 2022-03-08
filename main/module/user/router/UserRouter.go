package router

import (
	"github.com/gogf/gf/frame/g"
	"project/main/module/user/controller"
	"project/main/tool"
)

func init() {
	s := g.Server()
	tool.BindObjectReflect("/user", controller.UserController, s)
}
