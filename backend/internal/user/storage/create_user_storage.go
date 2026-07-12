package storage

import (
	"context"
	"time"
)

type userRecord struct {
	ID        string    `gorm:"column:id;primaryKey"`
	Username  string    `gorm:"column:username"`
	Password  string    `gorm:"column:password"`
	IsActive  bool      `gorm:"column:is_active;default:true"`
	CreatedAt time.Time `gorm:"column:created_at;autoCreateTime"`
	UpdatedAt time.Time `gorm:"column:updated_at;autoUpdateTime"`
}

func (userRecord) TableName() string {
	return "users"
}

func (s *userStorage) CreateUser(c context.Context, id, username, encryptedPassword string) error {
	timeNow := time.Now()

	user := userRecord{
		ID:        id,
		Username:  username,
		Password:  encryptedPassword,
		IsActive:  true,
		CreatedAt: timeNow,
		UpdatedAt: timeNow,
	}

	return s.DB.WithContext(c).Create(&user).Error
}
