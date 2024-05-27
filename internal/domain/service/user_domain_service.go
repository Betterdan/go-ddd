package service

import (
	"demo/internal/domain/repository"
)

type UserDomainService struct {
	userRepository repository.UserRepository
}

func NewUserDomainService(repo repository.UserRepository) *UserDomainService {
	return &UserDomainService{userRepository: repo}
}

func (u *UserDomainService) ChangeEmail(newEmail string) {
	// 业务规则，例如验证新邮箱
	return "lala"
}
