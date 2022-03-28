package controller

import (
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"project/main/module/user"
	"project/main/module/user/server"
	"project/main/tool"
	"project/main/tool/jwtTool"
)

func Register() {
	roleController := RoleController{}
	user.RegisterByStruct(roleController)
	s := g.Server("user").Group("/user")
	s.POST("/login", login)
	s.POST("/register", register)
	s = g.Server("user").Group("/role")
	s.POST("/create", createDefine)

}

func login(r *ghttp.Request) {
	password := r.Get("password", "").(string)
	name := r.Get("name", "").(string)
	value, err := server.Login(name, password)
	res := tool.Result{}
	if err == "" {
		data := map[string]any{
			"token": value,
		}
		r.Response.WriteJsonExit(res.Success("登录成功", &data))
	}
	r.Response.WriteJsonExit(res.Fail(500, err))
}

func register(r *ghttp.Request) {
	password := r.Get("password", "").(string)
	name := r.Get("name", "").(string)
	value, err := server.Register(name, password)
	res := tool.Result{}
	if err != "" {
		r.Response.WriteJsonExit(res.Fail(500, err))
	}
	data := map[string]any{
		"user":  value,
		"token": jwtTool.GetToken(value.UUID, nil),
	}
	r.Response.WriteJsonExit(res.Success("注册成功", &data))
}

func getUserInfo(r *ghttp.Request) {

}
