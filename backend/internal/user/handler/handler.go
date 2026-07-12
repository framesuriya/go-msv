package handler

import (
	"user/usecase"
)

func NewUserHandler(userUsecase usecase.UserUsecase) *userHandler {
	return &userHandler{
		UserUsecase: userUsecase,
	}
}

type userHandler struct {
	UserUsecase usecase.UserUsecase
}