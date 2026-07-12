package usecase

import (
	"context"
	"user/model"
	"user/storage"
)

type userUsecase struct {
	Storage storage.UserStorage
}

type UserUsecase interface {
	CreateUser(c context.Context, req model.CreateUserRequest) error
}

func NewUserUsecase(storage storage.UserStorage) UserUsecase {
	return &userUsecase{
		Storage: storage,
	}
}
