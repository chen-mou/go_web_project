package server

import (
	"google.golang.org/grpc"
)

var Sign chan byte
var S *grpc.Server

func init() {
	S = grpc.NewServer()
	Sign = make(chan byte, 1)
	Sign <- 1
}
