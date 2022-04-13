package server

import (
	"context"
	"github.com/micro-community/micro/v3/service/logger"
	pb "github.com/muskke/grpc-test/proto/v2"
	"net/http"
)

type Handler struct{}

func (h *Handler) Test(ctx context.Context, req *pb.Request, rsp *pb.Response) error {
	logger.Infof("Handler.Test - req: [%+v]", req)
	rsp.Version = pb.TestVersion_VERSION_2
	rsp.Code = http.StatusOK
	rsp.Data = make(map[string]string)
	rsp.Data["name"] = req.GetName()
	//rsp.Data["others"] = string(req.GetOthers())
	rsp.Height = req.GetHeight()
	rsp.Weight = req.GetWeight()
	logger.Infof("Handler.Test - rsp: [%+v]", rsp)
	return nil
}