package main

import (
	"github.com/micro-community/micro/v3/service"
	pbV2 "github.com/muskke/grpc-test/proto/v2"
	pbV3 "github.com/muskke/grpc-test/proto/v3"
	v2"github.com/muskke/grpc-test/server/v2"
	v3"github.com/muskke/grpc-test/server/v3"
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
	_ = pbV2.RegisterGrpcTestV2Handler(srv.Server(), &v2.Handler{})
	_ = pbV3.RegisterGrpcTestV3Handler(srv.Server(), &v3.Handler{})
	if err := srv.Run(); err != nil {
		panic(err)
	}
}
