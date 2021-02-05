package models

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Storage struct {
	*gorm.DB
}

func NewPostgresDB(dns string) (*Storage, error) {
	db, err := gorm.Open(postgres.Open(dns), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	err = db.AutoMigrate(&User{})
	if err != nil {
		return nil, err
	}

	return &Storage{db}, nil
}
