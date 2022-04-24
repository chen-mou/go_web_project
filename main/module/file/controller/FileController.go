package controller

import (
	"encoding/json"
	"github.com/chen-mou/gf/frame/g"
	"github.com/chen-mou/gf/net/ghttp"
	"project/main/module/file/entity"
	"project/main/module/file/middware"
	"project/main/module/file/server"
	"project/main/module/user"
	"project/main/tool"
	"strconv"
)

type FileController struct {
	uploadAvatar func(request *ghttp.Request) `role:"USER_BASE" path:"/file/avatar"`
}

func Register() {
	s := g.Server("user").Group("/file")
	s.Middleware(middware.Verify).POST("/upload", UploadBase)
	s.POST("/avatar", uploadAvatar)
	fileController := FileController{}
	user.RegisterByStruct(fileController)
}

func UploadBase(request *ghttp.Request) {
	jsonByte := request.GetFormString("file_attribute")
	var fileEntity entity.FileUpload
	json.Unmarshal([]byte(jsonByte), &fileEntity)
	tool.Analyse(fileEntity)
	file := request.GetUploadFile("file")
	index, _ := strconv.Atoi(request.GetString("index", 0))
	msg, err := server.Upload(index, file, fileEntity)
	if err != nil {
		request.Response.WriteJsonExit(tool.Result{}.Fail(500, err.Error()))
		return
	}
	request.Response.WriteJsonExit(tool.Result{}.Success(msg, nil))
}

func uploadAvatar(req *ghttp.Request) {

}

func get(req *ghttp.Request) {

}
