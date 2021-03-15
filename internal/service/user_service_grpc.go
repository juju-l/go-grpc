package service

import (
	pb "gitee.com/vipex/go-grpc/api/vipex.cc/oauth2/proto"
	"gitee.com/vipex/go-grpc/internal/domain/dao"
	// pri "gitee.com/vipex/go-grpc/internal/domain/interface"
	"gitee.com/vipex/go-grpc/internal/usecase"
)

func (rpc *UserServiceGrpc) Login(user pb.User) (pb.BaseResult, error) {
	var rst = pb.BaseResult{}
	if r, err := new(usecase.UserUseCase).Login(&dao.UserReq{user.User, user.Pswd}); err != nil {
		rst.Data = r
	}
	return rst, nil
}

type UserServiceGrpc struct {
	// pri.UnimplementedUserGrpcServer
}
