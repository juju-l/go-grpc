package usecase

import (
	"gitee.com/vipex/go-grpc/internal/domain/dao"
	pri "gitee.com/vipex/go-grpc/internal/domain/interface"
)

func (use *UserUseCase) Login(req *dao.UserReq) (bool, error) {
	return true, nil
}

type UserUseCase struct {
	userInterfaceUse pri.UserInterfaceUse
}
