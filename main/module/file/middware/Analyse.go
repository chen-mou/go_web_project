package middware

import (
	"github.com/chen-mou/gf/net/ghttp"
	"project/main/module/file"
	"project/main/tool"
)

func Verify(req *ghttp.Request) {
	ip := req.Header["Referer"][0]
	for x := range file.Allowed {
		if file.Allowed[x] == ip {
			req.Middleware.Next()
			return
		}
	}
	r := tool.Result{}
	req.Response.WriteJsonExit(r.Fail(403, "ip:"+ip+" 不在白名单内"))
}
