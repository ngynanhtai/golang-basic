package biz

import (
	"context"
	"go-demo/modules/item/model"
)

type GetItemStorage interface {
	GetItem(ctx context.Context, condition map[string]interface{}) (*model.TodoItem, error)
}

type GetItemBiz struct {
	store GetItemStorage
}

func NewGetItemBiz(store GetItemStorage) *GetItemBiz {
	return &GetItemBiz{store: store}
}

func (biz *GetItemBiz) GetItemById(ctx context.Context, id int) (*model.TodoItem, error) {
	data, err := biz.store.GetItem(ctx, map[string]interface{}{"id": id})
	if err != nil {
		return nil, err
	}

	return data, nil
}
