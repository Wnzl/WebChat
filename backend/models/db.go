package models

import (
	"github.com/cenkalti/backoff/v4"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Storage struct {
	*gorm.DB
}

func NewPostgresDB(dns string) (*Storage, error) {
	var db *gorm.DB

	err := backoff.Retry(func() (err error) {
		db, err = gorm.Open(postgres.Open(dns), &gorm.Config{})

		return
	}, backoff.WithMaxRetries(backoff.NewExponentialBackOff(), 5))

	if err != nil {
		return nil, err
	}

	err = db.AutoMigrate(&User{})
	if err != nil {
		return nil, err
	}

	return &Storage{db}, nil
}
