package v1

import (
	"context"
	v1_proto "gitee.com/vipex/go-grpc/api/vipex.cc/oauth2/v1/v1.proto"
	v1_dao "gitee.com/vipex/go-grpc/internal/domain/v1/v1.dao"
	v1_usecase "gitee.com/vipex/go-grpc/internal/usecase/v1"
)

func (rpc *UserGrpcHandler) Login(ctx context.Context, user *v1_proto.User, baseResult *v1_proto.BaseResult) error {
	// var rst = &v1_proto.BaseResult{}
	if r, err := new(v1_usecase.UserUseCase).Login(&v1_dao.UserReq{user.User, user.Pswd}); err == nil {
		/*rst*/baseResult.Data = r
	}
	/*baseResult = rst; */return nil
}

type UserGrpcHandler struct {
}
