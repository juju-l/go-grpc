package interfacepri

import "gitee.com/vipex/go-grpc/internal/domain/dao"

type UserInterfaceUse interface {
	Login(req *dao.UserReq) (bool, error)
}