package biz

import (
	"context"
	"errors"
	"go-demo/common"
	"go-demo/models/item"
	"strings"
)

type ItemStorage interface {
	GetItem(ctx context.Context, condition map[string]interface{}) (*item.TodoItem, error)
	ListItem(ctx context.Context, filter *item.Filter, paging *common.Paging, moreKeys ...string) ([]item.TodoItem, error)
	CreateItem(ctx context.Context, data *item.TodoItemCreation) error
	UpdateItem(ctx context.Context, condition map[string]interface{}, dataUpdate *item.TodoItemUpdate) error
	DeleteItem(ctx context.Context, condition map[string]interface{}) error
}

type ItemBiz struct {
	store ItemStorage
}

func NewItemBiz(store ItemStorage) *ItemBiz {
	return &ItemBiz{store: store}
}

func (biz *ItemBiz) GetItemById(ctx context.Context, id int) (*item.TodoItem, error) {
	data, err := biz.store.GetItem(ctx, map[string]interface{}{"id": id})
	if err != nil {
		return nil, common.ErrCannotGetEntity(item.EntityName, err)
	}

	return data, nil
}

func (biz *ItemBiz) ListItem(ctx context.Context, filter *item.Filter, paging *common.Paging) ([]item.TodoItem, error) {

	data, err := biz.store.ListItem(ctx, filter, paging)
	if err != nil {
		return nil, common.ErrCannotListEntity(item.EntityName, err)
	}

	return data, nil
}

func (biz *ItemBiz) CreateItem(ctx context.Context, data *item.TodoItemCreation) error {
	title := strings.TrimSpace(data.Title)

	if title == "" {
		return item.ErrTitleIsBlank
	}

	if err := biz.store.CreateItem(ctx, data); err != nil {
		return common.ErrCannotCreateEntity(item.EntityName, err)
	}

	return nil
}

func (biz *ItemBiz) UpdateItemById(ctx context.Context, id int, dataUpdate *item.TodoItemUpdate) error {
	data, err := biz.store.GetItem(ctx, map[string]interface{}{"id": id})

	if err != nil {
		if errors.Is(err, common.RecordNotFound) {
			return common.ErrCannotGetEntity(item.EntityName, err)
		}
		return common.ErrCannotUpdateEntity(item.EntityName, err)
	}

	if data.Status != nil && *data.Status == item.ItemStatusDeleted {
		return common.ErrEntityDeleted(item.EntityName, item.ErrItemIsDeleted)
	}

	if err := biz.store.UpdateItem(ctx, map[string]interface{}{"id": id}, dataUpdate); err != nil {
		return common.ErrCannotUpdateEntity(item.EntityName, err)
	}

	return nil
}

func (biz *ItemBiz) DeleteItemById(ctx context.Context, id int) error {
	data, err := biz.store.GetItem(ctx, map[string]interface{}{"id": id})

	if err != nil {
		if errors.Is(err, common.RecordNotFound) {
			return common.ErrCannotGetEntity(item.EntityName, err)
		}
		return common.ErrCannotDeleteEntity(item.EntityName, err)
	}

	if data.Status != nil && *data.Status == item.ItemStatusDeleted {
		return common.ErrEntityDeleted(item.EntityName, item.ErrItemIsDeleted)
	}

	if err := biz.store.DeleteItem(ctx, map[string]interface{}{"id": id}); err != nil {
		return common.ErrCannotDeleteEntity(item.EntityName, err)
	}

	return nil
}
