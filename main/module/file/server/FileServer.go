package server

import (
	"errors"
	"github.com/chen-mou/gf/frame/g"
	"github.com/chen-mou/gf/net/ghttp"
	"io/ioutil"
	"project/main/tool/encryption"

	"os"
	"project/main/module/file/entity"
	"project/main/module/file/model"
	"project/main/tool"
	"project/main/tool/dbTool"
	"strconv"
	"strings"
	"time"
)

var TempAddress = ""

var KeyPrefix = "FILE_UPLOAD"

func combo(upload *entity.FileUpload) (string, error) {
	fileEntity, ok, e := model.HasId(upload.FileId)
	if e != nil {
		return "", e
	}
	if !ok {
		return "", errors.New("文件不存在")
	}
	target, err1 := os.Create(fileEntity.Path)
	if err1 != nil {
		return "", err1
	}
	for i := range upload.Paths {
		file, err := os.Open(upload.Paths[i])
		if err != nil {
			return "", err
		}
		bt, _ := ioutil.ReadAll(file)
		target.Write(bt)
		file.Close()
		os.Remove(upload.Paths[i])
	}
	target.Close()
	return "COMPLETE", nil
}

func Upload(index int, file *ghttp.UploadFile, data entity.FileUpload) (string, error) {
	var fileUpload *entity.FileUpload
	dbTool.Get(KeyPrefix+"-"+data.FileId, fileUpload)
	if fileUpload == nil {
		return "", errors.New("文件不存在")
	}
	if g.Config().GetString("machine_name") == fileUpload.Machine {
		if index == fileUpload.Index-1 && index <= fileUpload.End {
			file.Filename = encryption.MD5Salt(strconv.FormatInt(time.Now().Unix(), 10), fileUpload.Machine) +
				strings.Split(file.Filename, ".")[0] +
				"-" + strconv.Itoa(index)
			_, err := file.Save(TempAddress)
			if err != nil {
				return "", err
			}
			fileUpload.Paths = append(fileUpload.Paths, TempAddress+"/"+file.Filename)
			if fileUpload.Index == fileUpload.End {
				fileUpload.Status = "COMPLETE"
				_, err = combo(fileUpload)
				if err != nil {
					return "", err
				}
				return "COMPLETE", nil
			}
			fileUpload.Index += 1
			dbTool.Set(KeyPrefix+"-"+data.FileId, fileUpload, time.Hour*3)
			return "SUCCESS", nil
		} else {
			return "", errors.New("上传有误")
		}
	} else {
		//重定向到别的服务器
	}
	return "", nil
}

func Create(fileData entity.File, uploader string, auth string, path string) (*entity.File, error) {
	fileData.Uploader = uploader
	fileData.Auth = auth
	fileData.MachineName = g.Config().GetString("machine_name")
	fileData.Status = "UPLOADING"
	fileData.FileId = getFileId(fileData)
	fileData.Path = path
	fileUpload := entity.FileUpload{
		Index: 0,
		/*计算有多少次*/
		End:     int(fileData.Size / 20480),
		FileId:  fileData.FileId,
		Machine: fileData.MachineName,
		Status:  fileData.Status,
	}
	file, err := model.Create(fileData, dbTool.Mysql)
	if err != nil {
		return nil, err
	}
	dbTool.Set(KeyPrefix+"-"+file.FileId, fileUpload, time.Hour*3)
	return file, nil
}

func sendToOtherServer(serverName string) {}

func getFileId(file entity.File) string {
	now := time.Now().Unix()
	value := file.Filename
	salt := strconv.FormatInt(now, 10) + dbTool.GetThreadID() + tool.Get("name")
	return encryption.MD5Salt(value, salt)
}
