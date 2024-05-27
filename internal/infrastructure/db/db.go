package db

import (
	"database/sql"
	"demo/internal/infrastructure/config"
	"fmt"
	_ "github.com/lib/pq"
)

func NewDB(config *config.Config) (*sql.DB, error) {
	dsn := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		config.DbConfig.DBHost, config.DbConfig.DBPort, config.DbConfig.DBUser, config.DbConfig.DBPassword, config.DbConfig.DBName,
	)
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}
