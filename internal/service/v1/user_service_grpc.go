package v1_service

//import (
//	v1_proto "gitee.com/vipex/go-grpc/api/vipex.cc/oauth2/v1/v1.proto"
//	"gitee.com/vipex/go-grpc/internal/domain/dao"
//	v1_interface "gitee.com/vipex/go-grpc/internal/domain/v1/v1.interface"
//	"gitee.com/vipex/go-grpc/internal/usecase"
//)
//
//func (rpc *UserServiceGrpc) Login(user v1_proto.User) (v1_proto.BaseResult, error) {
//	var rst = v1_proto.BaseResult{}
//	if r, err := new(usecase.UserUseCase).Login(&dao.UserReq{user.User, user.Pswd}); err != nil {
//		rst.Data = r
//	}
//	return rst, nil
//}
//
//type UserServiceGrpc struct {
//	v1_interface.UnimplementedUserGrpcServer
//}
