package biz

import (
	"context"
	"go-demo/common"
	"go-demo/modules/item/model"
	"strings"
)

type CreateItemStorage interface {
	CreateItem(ctx context.Context, data *model.TodoItemCreation) error
}

type CreateItemBiz struct {
	store CreateItemStorage
}

func NewCreateItemBiz(s CreateItemStorage) *CreateItemBiz {
	return &CreateItemBiz{s}
}

func (biz *CreateItemBiz) CreateItem(ctx context.Context, data *model.TodoItemCreation) error {
	title := strings.TrimSpace(data.Title)

	if title == "" {
		return model.ErrTitleIsBlank
	}

	if err := biz.store.CreateItem(ctx, data); err != nil {
		return common.ErrCannotCreateEntity(model.EntityName, err)
	}

	return nil
}
