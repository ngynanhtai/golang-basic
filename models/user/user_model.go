package user

import (
	"go-demo/common"
)

const (
	UserEntity = "User"
)

type User struct {
	common.SqlModel
	Email    string `json:"email" gorm:"column:email;unique"`
	Password string `json:"password" gorm:"column:password;"`
}

func (User) TableName() string { return "user" }

type UserCreation struct {
	Id       int    `json:"-" gorm:"column:id;"`
	Email    string `json:"email" gorm:"column:email;"`
	Password string `json:"password" gorm:"column:password;"`
}

func (UserCreation) TableName() string { return User{}.TableName() }
