package main

import (
	"flag"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	UserRouter "project/main/module/user/router"
)

var configMap = map[string]string{
	"dev":  "config/dev.toml",
	"dep":  "config/dep.toml",
	"test": "config/test.toml",
}

func register(server *ghttp.Server) {
	UserRouter.Register(server)
}

func main() {
	var env string
	flag.StringVar(&env, "env", "", "dev")
	flag.Parse()
	g.Cfg().SetFileName(configMap[env])
	s := g.Server()
	register(s)
	s.Run()
}
