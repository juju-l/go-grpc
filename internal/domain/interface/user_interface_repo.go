package interfacepri

import "gitee.com/vipex/go-grpc/internal/domain/dao"

type UserInterfaceRepo interface {
	Login(req *dao.UserReq) (bool, error)
}
