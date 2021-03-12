package v1

import (
	v1_interface "gitee.com/vipex/go-grpc/internal/domain/v1/v1.interface"
)

func (s *UserRepository) Login() error {
	return nil
}

type UserRepository struct {
	userRepository v1_interface.UserInterfaceRepo
}
