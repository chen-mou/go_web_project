package main

import (
	"flag"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
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
	s := g.Server()
	s.BindHandler("GET:/get", func(r *ghttp.Request) {
		r.Response.Writefln("success")
	})
	s.Run()
}
