package storage

import (
	"context"
	"go-demo/modules/item/model"
)

func (sql *SqlStore) UpdateItem(ctx context.Context, condition map[string]interface{}, dataUpdate *model.TodoItemUpdate) error {
	if err := sql.db.Where(condition).Updates(dataUpdate).Error; err != nil {
		return err
	}
	return nil
}
