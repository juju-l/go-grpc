package service

import (
	"context"
	pb "codeup.aliyun.com/vipex/go-grpc/api/vipex.cc/oauth2/proto"
	"codeup.aliyun.com/vipex/go-grpc/internal/domain/dao"
	"codeup.aliyun.com/vipex/go-grpc/internal/usecase"
)

func (rpc *UserServiceGrpc) Login(ctx context.Context, user *pb.User) (*pb.UserResult, error) {
	var rst = &pb.UserResult{}
	if r, err := new(usecase.UserUseCase).Login(&dao.UserReq{user.User, user.Pswd}); err == nil { // 
		rst.Data = r
	}
	return rst, nil
}

type UserServiceGrpc struct {
}
