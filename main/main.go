package main

import (
	"flag"
	"github.com/gogf/gf/frame/g"
	_ "project/main/module/user/controller"
)

var configMap = map[string]string{
	"dev":  "config/dev.toml",
	"dep":  "config/dep.toml",
	"test": "config/test.toml",
}

func main() {
	var env string
	flag.StringVar(&env, "env", "", "dev")
	flag.Parse()
	g.Cfg().SetFileName(configMap[env])
	s := g.Server("user")
	s.Run()
}
