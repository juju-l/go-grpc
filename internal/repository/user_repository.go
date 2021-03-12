package repository

import (
	interfacepri "gitee.com/vipex/go-grpc/internal/domain/interface"
)

func (s *UserRepository) Login() error {
	return nil
}

type UserRepository struct {
	userRepository interfacepri.UserInterfaceRepo
}
