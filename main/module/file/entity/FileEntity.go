package entity

import "project/main/tool/time"

type File struct {
	Id          int
	FileId      string
	Uploader    string
	MachineName string
	Filename    string `NotNull:"文件名不能为空"`
	Type        string `NotNull:"文件类型不能为空"`
	Suffix      string `NotNull:"文件后缀不能为空"`
	Path        string
	Auth        string
	Status      string
	Size        int64
	Ctime       time.Timestamp
	Mtime       time.Timestamp
}

type FileUpload struct {
	Id      int
	FileId  string
	Index   int
	End     int
	Machine string
	Paths   []string
	Status  string
	Ctime   time.Timestamp
	Mtime   time.Timestamp
}
