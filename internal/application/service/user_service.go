package service

import (
	"demo/internal/domain/repository"
)

type UserService struct {
	userRepository repository.UserRepository
}

func NewUserService(repo repository.UserRepository) *UserService {
	return &UserService{userRepository: repo}
}

//	func (s *UserAppService) RegisterUser(name, email string) (*model.User, error) {
//		user := &model.User{
//			ID:    generateID(),
//			Name:  name,
//			Email: email,
//		}
//		return user, s.userRepository.Save(user)
//	}
func (service *UserService) GetUser(id int64) (int, error) {
	return 1, nil
}
