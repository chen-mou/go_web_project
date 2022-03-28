package controller

import (
	"encoding/json"
	"github.com/chen-mou/gf/frame/g"
	"github.com/chen-mou/gf/net/ghttp"
	"project/main/module/user"
	"project/main/module/user/entity"
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

func createDefine(r *ghttp.Request) {
	res := tool.Result{}
	r.Response.WriteJsonExit(res.Success("test", nil))
}

func createBaseManager(r *ghttp.Request) {
	jsonByte := r.GetBody()
	var users []entity.User
	json.Unmarshal(jsonByte, &users)
}
