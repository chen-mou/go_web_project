package main

import (
	"flag"
	"github.com/gogf/gf/frame/g"
	userController "project/main/module/user/controller"
	"project/main/module/user/middware"
	_ "project/main/tool/dbTool"
)

var configMap = map[string]string{
	"dev":  "config/dev.yml",
	"dep":  "config/dep.yml",
	"test": "config/test.yml",
}

func register() {
	userController.Register()
}

func main() {
	var env string
	flag.StringVar(&env, "env", "", "dev")
	flag.Parse()
	g.Config().SetFileName(configMap[env])
	s := g.Server("user")
	s.Use(middware.CORS)
	s.Use(middware.JWT)
	register()
	s.Run()
}
