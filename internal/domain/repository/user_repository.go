package repository

import "demo/internal/domain/model"

/**
  接口设计、具体实现放 infrastructure/persistence
*/

type UserRepository interface {
	Create(user *model.User) error
	SaveInfo(user *model.User) error
	FindByID(id int64) (*model.User, error)
	DeleteUser(id int64) error
}
