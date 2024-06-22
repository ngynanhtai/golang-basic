package storage

import (
	"context"
	"go-demo/modules/item/model"
)

func (sql *SqlStore) DeleteItem(ctx context.Context, condition map[string]interface{}) error {
	deletedStatus := model.ItemStatusDeleted

	if err := sql.db.Table(model.TodoItem{}.TableName()).
		Where(condition).
		Updates(map[string]interface{}{
			"status": deletedStatus.String(),
		}).Error; err != nil {
		return err
	}
	return nil
}
