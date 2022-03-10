package main

import (
	"flag"
	"fmt"
	"github.com/gogf/gf/frame/g"
	userController "project/main/module/user/controller"
	_ "project/main/tool/dbTool"
)

var configMap = map[string]string{
	"dev":  "config/dev.toml",
	"dep":  "config/dep.toml",
	"test": "config/test.toml",
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
	fmt.Println(g.Config().GetFilePath())
	register()
	s.Run()
}
