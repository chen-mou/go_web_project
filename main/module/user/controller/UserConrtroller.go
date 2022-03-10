package controller

import (
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
)

func Register() {
	s := g.Server("user").Group("/user")
	s.POST("/login", login)
	s.POST("/register", register)

}

func login(r *ghttp.Request) {
	password := r.Get("password", "")
	name := r.Get("name", "")
	r.Response.Write(password)
	r.Response.Write(name)
}

func register(r *ghttp.Request) {

}
