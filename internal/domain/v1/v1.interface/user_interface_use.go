package v1_interface

import v1_dao "codeup.aliyun.com/vipex/go-grpc/internal/domain/v1/v1.dao"

type UserInterfaceUse interface {
	Login(req *v1_dao.UserLoginReq) (*v1_dao.User, error)
}
