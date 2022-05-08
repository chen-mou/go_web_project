package main

import (
	"flag"
	"github.com/chen-mou/gf/frame/g"
	"github.com/chen-mou/gf/net/ghttp"
	fileController "project/main/module/file/controller"
	userController "project/main/module/user/controller"
	"project/main/module/user/middware"
	"project/main/tool"
	_ "project/main/tool/dbTool"
	"project/main/tool/mq"
)

var configMap = map[string]string{
	"dev":  "config/dev.yml",
	"dep":  "config/dep.yml",
	"test": "config/test.yml",
}

func register() {
	userController.Register()
	fileController.Register()
}

func main() {
	var env string
	flag.StringVar(&env, "env", "", "dev")
	flag.Parse()
	g.Config().SetFileName(configMap[env])
	s := g.Server("user")
	s.SetClientMaxBodySize(20 * 1024 * 1024)
	s.Use(middware.CORS)
	s.Use(middware.JWT)
	s.Use(func(r *ghttp.Request) {
		defer func() {
			if err := recover(); err != any(nil) {
				r.Response.WriteJsonExit(tool.Result{}.Fail(500, err.(error).Error()))
			}
		}()
		r.Middleware.Next()
	})
	//mq.Register("test", nil)
	//mq.Ch <- true

	register()
	mq.Send("test", "message")
	s.Run()
}
