// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: grpc_test_v2.proto

package grpc_test

import (
	fmt "fmt"
	proto "google.golang.org/protobuf/proto"
	math "math"
)

import (
	context "context"
	api "github.com/micro-community/micro/v3/service/api"
	client "github.com/micro-community/micro/v3/service/client"
	server "github.com/micro-community/micro/v3/service/server"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// Reference imports to suppress errors if they are not otherwise used.
var _ api.Endpoint
var _ context.Context
var _ client.Option
var _ server.Option

// Api Endpoints for GrpcTestV2 service

func NewGrpcTestV2Endpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for GrpcTestV2 service

type GrpcTestV2Service interface {
	Test(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error)
}

type grpcTestV2Service struct {
	c    client.Client
	name string
}

func NewGrpcTestV2Service(name string, c client.Client) GrpcTestV2Service {
	return &grpcTestV2Service{
		c:    c,
		name: name,
	}
}

func (c *grpcTestV2Service) Test(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "GrpcTestV2.Test", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for GrpcTestV2 service

type GrpcTestV2Handler interface {
	Test(context.Context, *Request, *Response) error
}

func RegisterGrpcTestV2Handler(s server.Server, hdlr GrpcTestV2Handler, opts ...server.HandlerOption) error {
	type grpcTestV2 interface {
		Test(ctx context.Context, in *Request, out *Response) error
	}
	type GrpcTestV2 struct {
		grpcTestV2
	}
	h := &grpcTestV2Handler{hdlr}
	return s.Handle(s.NewHandler(&GrpcTestV2{h}, opts...))
}

type grpcTestV2Handler struct {
	GrpcTestV2Handler
}

func (h *grpcTestV2Handler) Test(ctx context.Context, in *Request, out *Response) error {
	return h.GrpcTestV2Handler.Test(ctx, in, out)
}
