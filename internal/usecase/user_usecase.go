package usecase

import (
	"codeup.aliyun.com/vipex/go-grpc/internal/domain/dao"
	pri "codeup.aliyun.com/vipex/go-grpc/internal/domain/interface"
	"codeup.aliyun.com/vipex/go-grpc/internal/repository"
)

func (use *UserUseCase) Login(req *dao.UserReq) (bool, error) {
	/***
	 * 该 userCase 层主要组装所有的业务逻辑
	 * 该单一业务直接返回，不表示该分层仅执行该内容，复杂业务时，在该服务层中，需继续其他实现
	***/
	return new(repository.UserRepository).Login(req)
}

type UserUseCase struct {
	userInterfaceUse pri.UserInterfaceUse
}
