package file

import (
	"context"
	"errors"
	"project/main/tool/rpg/client"
	"project/main/tool/rpg/server/file"
	"time"
)

func Upload() error {
	ctx, _ := context.WithTimeout(context.Background(), time.Second*5)
	val, ok := client.GetConnect(ctx)
	if !ok {
		panic("服务器繁忙")
	}
	c := file.NewFileServerClient(val)
	ctx, _ = context.WithTimeout(context.Background(), time.Second*5)
	js, _ := c.Upload(ctx, &file.File{})
	if js.GetCode() != 200 {
		return errors.New(js.GetMsg())
	}
	return nil
}
