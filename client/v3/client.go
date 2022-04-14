package main

import (
	"context"
	grpcCli "github.com/micro-community/micro/v3/service/client/grpc"
	"github.com/micro-community/micro/v3/service/logger"
	pb "github.com/muskke/grpc-test/proto/v3"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"sync"
	"time"

	"github.com/micro-community/micro/v3/service/client"
)

const addr = "localhost:10000"

var instance pb.GrpcTestV3Service
var once sync.Once

// GetInstance 获取服务单例对象
func GetInstance() pb.GrpcTestV3Service {
	once.Do(func() {
		cli := grpcCli.NewClient(client.Proxy("localhost:10000"))
		instance = pb.NewGrpcTestV3Service("grpc-test-service", cli)
	})
	return instance
}

func GrpcClient() pb.GrpcTestV3Client {
	conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		logger.Fatalf("did not connect: %v", err)
	}
	//defer conn.Close()

	return pb.NewGrpcTestV3Client(conn)
}

func DoTestV3(ctx context.Context, req *pb.Request) {
	logger.Infof("DoTestV3 req: [%+v]", req)

	rsp, err := GetInstance().Test(ctx, req)
	if err != nil {
		logger.Errorf("DoTestV3 failed with err:%v", req, err)
		return
	}
	logger.Infof("DoTestV3 rsp: [%+v]", rsp)
	return
}

func main() {
	req := &pb.Request{
		Name: "Test",
		Age:  100,
		//Others: []byte("height: 180cm,weight: 65.0kg"),
		Height: 180.50,
		Weight: 65.0,
	}
	ctx, cancel := context.WithTimeout(context.TODO(), time.Second*1000000)
	defer cancel()
	DoTestV3(ctx,req)

	// 原生GRPC调用
	rsp, err := GrpcClient().Test(ctx, req)
	if err != nil {
		logger.Errorf("GrpcClient.Test failed: %v", err)
		return
	}
	logger.Infof("GrpcClient.Test rsp: [%+v]", rsp)
}
