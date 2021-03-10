package handler

import (
	base "gitee.com/vipex/go-grpc/gen/base"
	user "gitee.com/vipex/go-grpc/gen/business/user"
	"golang.org/x/net/context"
)

func (s *Server) Login(ctx context.Context, in *user.User) (*base.BaseResult, error) {
	var result = &base.BaseResult{}
	result.Result = &base.Result{ ErrCode: 0, ErrMsg: "ok" }
	if in.User == "admin" && in.Pswd == "654321" {
		result.Data.TypeUrl = "@bool"
		result.Data.Value = []byte("true")
	}
	result.ErrorDes = &base.ErrorDes{ UserMsg: "", CodeMsg: "" }
	return result, nil
}

//
type Server struct {
	user.UnimplementedUserGrpcServer
}

