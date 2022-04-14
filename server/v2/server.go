package server

import (
	"context"
	"net/http"

	"github.com/micro-community/micro/v3/service/logger"
	pb "github.com/muskke/grpc-test/proto/v2"
)

type Handler struct{}

func (h *Handler) Test(ctx context.Context, req *pb.Request, rsp *pb.Response) error {
	logger.Infof("Handler.Test - req: [%+v]", req)
	rsp.Version = pb.TestVersion_VERSION_2
	rsp.Code = http.StatusOK
	rsp.Data = make(map[string]string)
	rsp.Data["name"] = req.GetName()
	rsp.Data["others"] = string(req.GetOthers())
	rsp.Height = 1.1
	rsp.Weight = req.GetWeight()
	logger.Infof("Handler.Test - rsp: [%+v]", rsp)
	return nil
}
