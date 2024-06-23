package storage

import (
	"context"
	"go-demo/common"
	"go-demo/modules/item/model"
)

func (sql *SqlStore) CreateItem(ctx context.Context, data *model.TodoItemCreation) error {
	if err := sql.db.Create(&data).Error; err != nil {
		return common.ErrDB(err)
	}
	return nil
}
