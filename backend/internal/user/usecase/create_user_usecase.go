package usecase

import (
	"context"
	"user/model"

	"github.com/google/uuid"
)

func (u *userUsecase) CreateUser(c context.Context, req model.CreateUserRequest) error {

	userID, err := newUserID()
	if err != nil {
		return err
	}

	if err := u.Storage.CreateUser(c, userID, req.Username, req.Password); err != nil {
		return err
	}
	return nil
}

func newUserID() (string, error) {
	return uuid.New().String(), nil

}
