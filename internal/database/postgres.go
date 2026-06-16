package database

import (
	"fmt"

	"github.com/priyanshu334/gopr1/internal/config"
	"github.com/priyanshu334/gopr1/internal/task"
	"github.com/priyanshu334/gopr1/internal/user"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Connect(cfg *config.Config) (*gorm.DB, error) {
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		cfg.DBHost,
		cfg.DBUser,
		cfg.DBPassword,
		cfg.DBName,
		cfg.DBPort,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	err = db.AutoMigrate(
		&user.User{},
		&task.Task{},
	)
	if err != nil {
		return nil, err
	}

	return db, nil
}
