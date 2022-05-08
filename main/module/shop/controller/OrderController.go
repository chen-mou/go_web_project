package controller

import "github.com/chen-mou/gf/net/ghttp"

type orderController struct {
	order func(*ghttp.Request) `path:"" role:"USER_BASE"`
}

func order(req *ghttp.Request) {}
