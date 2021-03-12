package service

import (
	pb "gitee.com/vipex/go-grpc/api/vipex.cc/oauth2/v1/v1.proto"
	"gitee.com/vipex/go-grpc/internal/domain/dao"
	"gitee.com/vipex/go-grpc/internal/usecase"
)

func (rpc *UserGrpcHandler) Login(user pb.User) (pb.BaseResult, error) {
	var rst = pb.BaseResult{}
	if r, err := new(usecase.UserUseCase).Login(&dao.UserReq{user.User, user.Pswd}); err != nil {
		rst.Data = r
	}
	return rst, nil
}

type UserGrpcHandler struct {
}
