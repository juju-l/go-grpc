package v1_usecase

import (
	v1_dao "gitee.com/vipex/go-grpc/internal/domain/v1/v1.dao"
	v1_interface "gitee.com/vipex/go-grpc/internal/domain/v1/v1.interface"
)

func (use *UserUseCase) Login(req *v1_dao.UserReq) (bool, error) {
	return true, nil
}

type UserUseCase struct {
	userInterfaceUse v1_interface.UserInterfaceUse
}
