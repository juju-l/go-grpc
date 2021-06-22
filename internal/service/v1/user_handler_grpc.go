package v1_service

import (
	"context"
	v1_proto "codeup.aliyun.com/vipex/go-grpc/api/vipex.cc/oauth2/v1/v1.proto"
	v1_dao "codeup.aliyun.com/vipex/go-grpc/internal/domain/v1/v1.dao"
	v1_usecase "codeup.aliyun.com/vipex/go-grpc/internal/usecase/v1"
)

func (hdl *UserGrpcHandler) Login(ctx context.Context, user *v1_proto.UserLoginReq, userResult *v1_proto.UserResult) error {
	// var rst = &v1_proto.UserResult{}
	if r, err := new(v1_usecase.UserUseCase).Login(&v1_dao.UserLoginReq{user.User, user.Pswd}); err == nil {
		/*rst*/userResult.Data = &v1_proto.User{ User: r.User }
	}
	/*userResult = rst; */return nil
}

type UserGrpcHandler struct {
}
