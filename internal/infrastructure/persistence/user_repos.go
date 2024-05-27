package persistence

import (
	"database/sql"
	"demo/internal/domain/model"
	"demo/internal/domain/repository"
)

type UserRepositoryImpl struct {
	DB *sql.DB
}

func NewUserRepository(db *sql.DB) repository.UserRepository {
	return &UserRepositoryImpl{DB: db}
}

func (r *UserRepositoryImpl) FindByID(id int64) (*model.User, error) {
	// Here we mock a user; in a real implementation, this would query a database
	return &model.User{ID: id, Name: "John Doe"}, nil
}
