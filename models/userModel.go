package models

import (
	"errors"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Email    string `json:"email" gorm:"column:email;type:varchar(100);unique"`
	Password string `json:"password" gorm:"column:password;type:varchar(30)"`
}

var (
	ErrCannotSyncUser = errors.New("cannot sync user entity")
)
