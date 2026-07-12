package storage

import (
	"context"

	"gorm.io/gorm"
)

func NewUserStorage(db *gorm.DB) UserStorage {
	return &userStorage{
		DB: db,
	}
}

type userStorage struct {
	DB *gorm.DB
}

type UserStorage interface {
	// Define methods for user storage operations here
	CreateUser(c context.Context, id, username, encryptedPassword string) error
}
