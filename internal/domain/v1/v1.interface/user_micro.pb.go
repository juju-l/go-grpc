// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: user.proto

package v1_interface

import v1_proto "gitee.com/vipex/go-grpc/api/vipex.cc/oauth2/v1/v1.proto"

import (
	fmt "fmt"
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

// Api Endpoints for UserGrpc service

func NewUserGrpcEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for UserGrpc service

type UserGrpcService interface {
	Login(ctx context.Context, in *v1_proto.User, opts ...client.CallOption) (*v1_proto.BaseResult, error)
}

type userGrpcService struct {
	c    client.Client
	name string
}

func NewUserGrpcService(name string, c client.Client) UserGrpcService {
	return &userGrpcService{
		c:    c,
		name: name,
	}
}

func (c *userGrpcService) Login(ctx context.Context, in *v1_proto.User, opts ...client.CallOption) (*v1_proto.BaseResult, error) {
	req := c.c.NewRequest(c.name, "UserGrpc.Login", in)
	out := new(v1_proto.BaseResult)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for UserGrpc service

type UserGrpcHandler interface {
	Login(context.Context, *v1_proto.User, *v1_proto.BaseResult) error
}

func RegisterUserGrpcHandler(s server.Server, hdlr UserGrpcHandler, opts ...server.HandlerOption) error {
	type userGrpc interface {
		Login(ctx context.Context, in *v1_proto.User, out *v1_proto.BaseResult) error
	}
	type UserGrpc struct {
		userGrpc
	}
	h := &userGrpcHandler{hdlr}
	return s.Handle(s.NewHandler(&UserGrpc{h}, opts...))
}

type userGrpcHandler struct {
	UserGrpcHandler
}

func (h *userGrpcHandler) Login(ctx context.Context, in *v1_proto.User, out *v1_proto.BaseResult) error {
	return h.UserGrpcHandler.Login(ctx, in, out)
}
