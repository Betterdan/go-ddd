package db

import (
	"demo/internal/infrastructure/config"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"time"
)

func NewDB(config *config.Config) (*gorm.DB, error) {
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		config.DbConfig.DBUser, config.DbConfig.DBPassword,
		config.DbConfig.DBHost, config.DbConfig.DBPort, config.DbConfig.DBName,
	)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	// 获取通用数据库对象 sql.DB ，以便使用其提供的函数
	sqlDB, err := db.DB()
	if err != nil {
		return nil, err
	}

	// 设置最大空闲连接数
	sqlDB.SetMaxIdleConns(10)

	// 设置最大打开连接数
	sqlDB.SetMaxOpenConns(100)

	// 设置了连接可复用的最大时间
	sqlDB.SetConnMaxLifetime(time.Hour)

	return db, nil
}
