package biz

import (
	"context"
	"go-demo/common"
	"go-demo/modules/item/model"
)

type ListItemStorage interface {
	ListItem(
		ctx context.Context,
		filter *model.Filter,
		paging *common.Paging,
		moreKeys ...string,
	) ([]model.TodoItem, error)
}

type ListItemBiz struct {
	store ListItemStorage
}

func NewListItemBiz(store ListItemStorage) *ListItemBiz {
	return &ListItemBiz{store: store}
}

func (biz *ListItemBiz) ListItem(
	ctx context.Context,
	filter *model.Filter,
	paging *common.Paging,
) ([]model.TodoItem, error) {

	data, err := biz.store.ListItem(ctx, filter, paging)
	if err != nil {
		return nil, common.ErrCannotListEntity(model.EntityName, err)
	}

	return data, nil
}
