package repository

import "demo/internal/domain/model"

type UserRepository interface {
	Save(user *model.User) error
	FindByID(id string) (*model.User, error)
}
