package main

import (
	"github.com/micro-community/micro/v3/service"
	pb "github.com/muskke/grpc-test/proto/v2"
	"github.com/muskke/grpc-test/server"
)

const (
	Addr    = "0.0.0.0:10000"
	Version = "v0.0.1"
)

func main() {
	srv := &service.Service{}
	srv.Init(
		service.Name("grpc-test-server"),
		service.Address(Addr),
		service.Version(Version),
	)
	_ = pb.RegisterGrpcTestHandler(srv.Server(), &server.Handler{})
	if err := srv.Run(); err != nil {
		panic(err)
	}
}
