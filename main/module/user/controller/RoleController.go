package controller

import (
	"encoding/json"
	"github.com/chen-mou/gf/frame/g"
	"github.com/chen-mou/gf/net/ghttp"
	"project/main/module/user"
	"project/main/module/user/server"
	"project/main/tool"
)

type RoleController struct {
	create func(r *ghttp.Request) `path:"/role/create" role:"ADMIN"`
}

func roleRegister() {
	s := g.Server("user").Group("/role")
	s.POST("/create", createDefine)
	roleController := RoleController{}
	user.RegisterByStruct(roleController)
}

func createRole(r *ghttp.Request) {}

func createDefine(r *ghttp.Request) {
	jsonByte := r.GetBody()
	res := tool.Result{}
	var users []string
	json.Unmarshal(jsonByte, &users)
	roles, err := server.CreateBaseManager(users, nil, 6)
	if err != nil {
		r.Response.WriteJsonExit(res.Fail(500, err.Error()))
		return
	}
	r.Response.WriteJsonExit(res.Success("创建成功", &map[string]any{
		"data": roles,
	}))
}
