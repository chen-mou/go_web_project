package middware

import (
	"github.com/gogf/gf/net/ghttp"
	"project/main/module/user"
	"project/main/module/user/entity"
	"project/main/module/user/model"
	"project/main/module/user/server"
	"project/main/tool"
	"project/main/tool/jwtTool"
)

func JWT(r *ghttp.Request) {
	targetRole := user.Policies[r.Request.RequestURI]
	if targetRole == nil {
		r.Middleware.Next()
		return
	}
	token := r.Header.Get("X-Auth-Token")
	claim, err := jwtTool.ParseToken(token)
	if err != nil {
		r.Response.WriteJsonExit(tool.Result{}.Fail(500, err.Error()))
		return
	}
	userRoles, errStr := server.GetUserRoleByUUID(claim.UUID)
	if errStr != "" || userRoles == nil {
		r.Response.WriteJsonExit(tool.Result{}.Fail(500, errStr))
		return
	}
	for i := range userRoles {
		userRole := model.Mix(userRoles[i].Role)
		if model.Verify(*userRole, *targetRole) {
			r.Attribute["CurrentUser"] = entity.User{
				Roles: userRoles,
				UUID:  claim.UUID,
			}
			r.Middleware.Next()
			return
		}
	}
	r.Response.WriteJsonExit(tool.Result{}.Fail(403, "没有足够的权限"))

}

func CORS(r *ghttp.Request) {
	r.Response.CORSDefault()
	r.Middleware.Next()
}
