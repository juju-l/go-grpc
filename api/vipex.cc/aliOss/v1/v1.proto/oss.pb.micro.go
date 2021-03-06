// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: api/vipex.cc/aliOss/v1/v1.proto/oss.proto

package v1_proto

import (
	fmt "fmt"
	_ "codeup.aliyun.com/vipex/go-grpc/api/vipex.cc/public/v1/v1.proto"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

import (
	context "context"
	api "github.com/micro/go-micro/v2/api"
	client "github.com/micro/go-micro/v2/client"
	server "github.com/micro/go-micro/v2/server"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

// Reference imports to suppress errors if they are not otherwise used.
var _ api.Endpoint
var _ context.Context
var _ client.Option
var _ server.Option

// Api Endpoints for OssGrpc service

func NewOssGrpcEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for OssGrpc service

type OssGrpcService interface {
	Test(ctx context.Context, in *Oss, opts ...client.CallOption) (*OssResult, error)
}

type ossGrpcService struct {
	c    client.Client
	name string
}

func NewOssGrpcService(name string, c client.Client) OssGrpcService {
	return &ossGrpcService{
		c:    c,
		name: name,
	}
}

func (c *ossGrpcService) Test(ctx context.Context, in *Oss, opts ...client.CallOption) (*OssResult, error) {
	req := c.c.NewRequest(c.name, "OssGrpc.Test", in)
	out := new(OssResult)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for OssGrpc service

type OssGrpcHandler interface {
	Test(context.Context, *Oss, *OssResult) error
}

func RegisterOssGrpcHandler(s server.Server, hdlr OssGrpcHandler, opts ...server.HandlerOption) error {
	type ossGrpc interface {
		Test(ctx context.Context, in *Oss, out *OssResult) error
	}
	type OssGrpc struct {
		ossGrpc
	}
	h := &ossGrpcHandler{hdlr}
	return s.Handle(s.NewHandler(&OssGrpc{h}, opts...))
}

type ossGrpcHandler struct {
	OssGrpcHandler
}

func (h *ossGrpcHandler) Test(ctx context.Context, in *Oss, out *OssResult) error {
	return h.OssGrpcHandler.Test(ctx, in, out)
}
