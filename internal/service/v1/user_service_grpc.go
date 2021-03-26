package v1_service

import (
	"context"
	v1_proto "gitee.com/vipex/go-grpc/api/vipex.cc/oauth2/v1/v1.proto"
	v1_dao "gitee.com/vipex/go-grpc/internal/domain/v1/v1.dao"
	v1_usecase "gitee.com/vipex/go-grpc/internal/usecase/v1"
)

func (rpc *UserServiceGrpc) Login(ctx context.Context, user *v1_proto.UserLoginReq) (*v1_proto.UserResult, error) {
	var rst = &v1_proto.UserResult{}
	if r, err := new(v1_usecase.UserUseCase).Login(&v1_dao.UserLoginReq{user.User, user.Pswd}); err == nil { //
		rst.Data = &v1_proto.User{ User: r.User }
	}
	return rst, nil
}

type UserServiceGrpc struct {
}
