package v1_interface

import v1_dao "gitee.com/vipex/go-grpc/internal/domain/v1/v1.dao"

type UserInterfaceRepo interface {
	Login(req *v1_dao.UserReq) (bool, error)
}
