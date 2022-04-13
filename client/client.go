package main

import (
	"context"
	grpcCli "github.com/micro-community/micro/v3/service/client/grpc"
	"github.com/micro-community/micro/v3/service/logger"
	pb "github.com/muskke/grpc-test/proto/v1"
	"sync"
	"time"

	"github.com/micro-community/micro/v3/service/client"
)

var instance pb.GrpcTestService
var once sync.Once

// GetInstance 获取服务单例对象
func GetInstance() pb.GrpcTestService {
	once.Do(func() {
		cli := grpcCli.NewClient(client.Proxy("localhost:10000"))
		instance = pb.NewGrpcTestService("grpc-test-service", cli)
	})
	return instance
}

func DoTestV1(ctx context.Context) {
	req := &pb.Request{
		Name:   "Test",
		Age:    100,
		Others: []byte("height: 180cm,weight: 65.0kg"),
	}
	logger.Infof("DoTestV1 req: [%+v]", req)

	rsp, err := GetInstance().Test(ctx, req)
	if err != nil {
		logger.Errorf("DoTestV1 failed with err:%v", req, err)
		return
	}
	logger.Infof("DoTestV1 rsp: [%+v]", rsp)
	return
}

func main() {
	ctx, cancel := context.WithTimeout(context.TODO(), time.Second)
	defer cancel()
	DoTestV1(ctx)
}
