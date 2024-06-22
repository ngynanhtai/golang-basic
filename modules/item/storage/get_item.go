package storage

import (
	"context"
	"go-demo/modules/item/model"
)

func (sql *SqlStore) GetItem(ctx context.Context, condition map[string]interface{}) (*model.TodoItem, error) {
	var data model.TodoItem
	if err := sql.db.Where(condition).First(&data).Error; err != nil {
		return nil, err
	}
	return &data, nil
}
