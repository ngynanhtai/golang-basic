package storage

import (
	"context"
	"errors"
	"go-demo/common"
	"go-demo/modules/item/model"
	"gorm.io/gorm"
)

func (sql *SqlStore) GetItem(ctx context.Context, condition map[string]interface{}) (*model.TodoItem, error) {
	var data model.TodoItem
	if err := sql.db.Where(condition).First(&data).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, common.RecordNotFound
		}
		return nil, common.ErrDB(err)
	}
	return &data, nil
}
