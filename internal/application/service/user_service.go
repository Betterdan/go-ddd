package service

import (
	"demo/internal/domain/model"
	"demo/internal/domain/repository"
	"demo/internal/domain/service"
)

type UserAppService struct {
	userRepository    repository.UserRepository
	userDomainService service.UserDomainService
}

func NewUserAppService(repo repository.UserRepository) *UserAppService {
	return &UserAppService{userRepository: repo}
}

func (s *UserAppService) RegisterUser(name, email string) (*model.User, error) {
	user := &model.User{
		ID:    generateID(),
		Name:  name,
		Email: email,
	}
	return user, s.userRepository.Save(user)
}
