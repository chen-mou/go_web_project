package controller

import (
	"encoding/json"
	"github.com/gogf/gf/net/ghttp"
	"project/main/module/user/entity"
	"project/main/tool"
)

type RoleController struct {
	create func(r *ghttp.Request) `path:"/role/create" role:"ADMIN"`
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
