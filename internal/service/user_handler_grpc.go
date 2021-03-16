package service

import (
	"context"
	pb "gitee.com/vipex/go-grpc/api/vipex.cc/oauth2/proto"
	"gitee.com/vipex/go-grpc/internal/domain/dao"
	"gitee.com/vipex/go-grpc/internal/usecase"
)

func (rpc *UserGrpcHandler) Login(ctx context.Context, user *pb.User, baseResult *pb.BaseResult) error {
	// var rst = &pb.BaseResult{}
	if r, err := new(usecase.UserUseCase).Login(&dao.UserReq{user.User, user.Pswd}); err == nil {
		/*rst*/baseResult.Data = r
	}
	/*baseResult = rst; */return nil
}

type UserGrpcHandler struct {
}
