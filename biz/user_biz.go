package biz

import (
	"context"
	"errors"
	"fmt"
	"go-demo/common"
	"go-demo/models/user"
	"golang.org/x/crypto/bcrypt"
)

type UserStorage interface {
	CreateUser(ctx context.Context, data *user.UserCreation) error
	GetUser(ctx context.Context, condition map[string]interface{}) (*user.User, error)
}

type UserBiz struct {
	store UserStorage
}

func NewUserBiz(store UserStorage) *UserBiz {
	return &UserBiz{store}
}

func (biz *UserBiz) CreateUser(ctx context.Context, data *user.UserCreation) error {
	val, err := bcrypt.GenerateFromPassword([]byte(data.Password), bcrypt.MinCost)
	if err != nil {
		return common.ErrInvalidRequest(err)
	}

	data.Password = string(val)
	if err := biz.store.CreateUser(ctx, data); err != nil {
		return common.ErrCannotCreateEntity(user.UserEntity, err)
	}
	return nil
}

func (biz *UserBiz) GetUserByEmail(ctx context.Context, email string) (*user.User, error) {
	if email == "" {
		return nil, common.ErrInvalidRequest(errors.New(fmt.Sprintf(common.ErrIsBlank, "email")))
	}
	data, err := biz.store.GetUser(ctx, map[string]interface{}{"email": email})
	if err != nil {
		return nil, common.ErrCannotGetEntity(user.UserEntity, err)
	}
	return data, nil
}

func (biz *UserBiz) GetUserById(ctx context.Context, id int) (*user.User, error) {
	data, err := biz.store.GetUser(ctx, map[string]interface{}{"id": id})
	if err != nil {
		return nil, common.ErrCannotGetEntity(user.UserEntity, err)
	}
	return data, nil
}
