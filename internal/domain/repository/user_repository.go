package repository

import "demo/internal/domain/model"

/**
  接口设计、具体实现放 infrastructure/persistence
*/

type UserRepository interface {
	//Save(user *model.User) error
	FindByID(id int64) (*model.User, error)
}
