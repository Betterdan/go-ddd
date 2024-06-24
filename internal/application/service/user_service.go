package service

import (
	"context"
	"demo/internal/application/dto"
	"demo/internal/domain/model"
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
func (service *UserService) GetUser(ctx context.Context, id int64) (*dto.UserDTO, error) {
	user, err := service.userRepository.FindByID(id)
	if err != nil {
		return nil, err
	}
	return toUserDTO(user), nil
}

func toUserDTO(user *model.User) *dto.UserDTO {
	if user == nil {
		return nil
	}
	return &dto.UserDTO{
		ID:    user.ID,
		Name:  user.Name,
		Email: user.Email,
	}
}
