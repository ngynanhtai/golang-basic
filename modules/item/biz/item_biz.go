package biz

import (
	"context"
	"go-demo/common"
	"go-demo/modules/item/model"
	"strings"
)

type ItemStorage interface {
	ListItem(ctx context.Context, filter *model.Filter, paging *common.Paging, moreKeys ...string) ([]model.TodoItem, error)
	GetItem(ctx context.Context, condition map[string]interface{}) (*model.TodoItem, error)
	CreateItem(ctx context.Context, data *model.TodoItemCreation) error
	UpdateItem(ctx context.Context, condition map[string]interface{}, dataUpdate *model.TodoItemUpdate) error
	DeleteItem(ctx context.Context, condition map[string]interface{}) error
}

type ItemBiz struct {
	store ItemStorage
}

func NewItemBiz(store ItemStorage) *ItemBiz {
	return &ItemBiz{store: store}
}

func (biz *ItemBiz) ListItem(ctx context.Context, filter *model.Filter, paging *common.Paging) ([]model.TodoItem, error) {

	data, err := biz.store.ListItem(ctx, filter, paging)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (biz *ItemBiz) GetItemById(ctx context.Context, id int) (*model.TodoItem, error) {
	data, err := biz.store.GetItem(ctx, map[string]interface{}{"id": id})
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (biz *ItemBiz) CreateItem(ctx context.Context, data *model.TodoItemCreation) error {
	title := strings.TrimSpace(data.Title)

	if title == "" {
		return model.ErrTitleIsBlank
	}

	if err := biz.store.CreateItem(ctx, data); err != nil {
		return err
	}

	return nil
}

func (biz *ItemBiz) UpdateItemById(ctx context.Context, id int, dataUpdate *model.TodoItemUpdate) error {
	data, err := biz.store.GetItem(ctx, map[string]interface{}{"id": id})

	if err != nil {
		return err
	}

	if data.Status != nil && *data.Status == model.ItemStatusDeleted {
		return model.ErrItemIsDeleted
	}

	if err := biz.store.UpdateItem(ctx, map[string]interface{}{"id": id}, dataUpdate); err != nil {
		return err
	}

	return nil
}

func (biz *ItemBiz) DeleteItemById(ctx context.Context, id int) error {
	data, err := biz.store.GetItem(ctx, map[string]interface{}{"id": id})

	if err != nil {
		return err
	}

	if data.Status != nil && *data.Status == model.ItemStatusDeleted {
		return model.ErrItemIsDeleted
	}

	if err := biz.store.DeleteItem(ctx, map[string]interface{}{"id": id}); err != nil {
		return err
	}

	return nil
}
