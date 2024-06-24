package persistence

import (
	"demo/internal/domain/model"
	"demo/internal/domain/repository"
	"gorm.io/gorm"
)

type UserRepositoryImpl struct {
	DB *gorm.DB
}

func NewUserRepository(db *gorm.DB) repository.UserRepository {
	return &UserRepositoryImpl{DB: db}
}

func (r *UserRepositoryImpl) FindByID(id int64) (*model.User, error) {
	var user model.User
	err := r.DB.First(&user, id).Error
	return &user, err
}

func (r *UserRepositoryImpl) SaveInfo(user *model.User) error {
	return r.DB.Save(user).Error
}

func (r *UserRepositoryImpl) Create(user *model.User) error {
	return r.DB.Create(user).Error
}

func (r *UserRepositoryImpl) DeleteUser(id int64) error {
	return r.DB.Delete(&model.User{}, id).Error
}
