package handler

import (
	// base "gitee.com/vipex/go-grpc/gen/base"
	user "gitee.com/vipex/go-grpc/gen/business/user"
	"github.com/micro/go-micro/v2/server"
	"golang.org/x/net/context"
)

func RegisterUser(s server.Server) {
	user.RegisterUserGrpcHandler(s, *new(user.UserGrpcHandler))
}

func (s *UserHandler) Login(ctx context.Context, in *user.User) (* /*base*/ user.BaseResult, error) {
	var result = & /*base*/ user.BaseResult{}
	result.Result = & /*base*/ user.Result{ErrCode: 0, ErrMsg: "ok"}
	if in.User == "admin" && in.Pswd == "654321" {
		result.Data = true
	}
	result.ErrorDes = & /*base*/ user.ErrorDes{UserMsg: "", CodeMsg: ""}
	return result, nil
}

//
type UserHandler struct {
	// user.UnimplementedUserGrpcHandler
}
