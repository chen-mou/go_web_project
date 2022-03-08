package controller

import (
	"github.com/gogf/gf/net/ghttp"
)

type userController struct {
	Login    func(request *ghttp.Request) `method:"POST" address:"/login"`
	Register func(request *ghttp.Request) `method:"GET" address:"/register"`
}

var UserController *userController

func init() {
	UserController = new(userController)
	UserController.Login = login
	UserController.Register = register
}

func login(r *ghttp.Request) {
	password := r.Get("password", "")
	name := r.Get("name", "")
	r.Response.Write(password)
	r.Response.Write(name)
}

func register(r *ghttp.Request) {

}
