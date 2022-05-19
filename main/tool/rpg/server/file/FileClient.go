package file

import (
	"context"
	"project/main/tool/rpg/server"
)

type Server struct{}

func (s *Server) mustEmbedUnimplementedFileServerServer() {
	//TODO implement me
	panic("implement me")
}

func init() {
	select {
	case <-server.Sign:
		RegisterFileServerServer(server.S, &Server{})
	}
}

func (s *Server) Upload(ctx context.Context, in *File) (*JSON, error) {
	return &JSON{
		Code: 200,
		Msg:  "success",
	}, nil
}
