package storage

import (
	"context"
	"errors"
	"go-demo/common"
	"go-demo/models/item"
	"gorm.io/gorm"
)

func (sql *SqlStore) GetItem(ctx context.Context, condition map[string]interface{}) (*item.TodoItem, error) {
	var data item.TodoItem
	if err := sql.db.Where(condition).First(&data).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, common.RecordNotFound
		}
		return nil, common.ErrDB(err)
	}
	return &data, nil
}

func (sql *SqlStore) ListItem(ctx context.Context, filter *item.Filter, paging *common.Paging, moreKeys ...string) ([]item.TodoItem, error) {
	var result []item.TodoItem

	db := sql.db.Where("status <> ?", "Deleted")

	if f := filter; f != nil {
		if v := f.Status; v != "" {
			db = db.Where("status = ?", v)
		}
	}

	if err := db.Table(item.TodoItem{}.TableName()).Count(&paging.Total).Error; err != nil {
		return nil, common.ErrDB(err)
	}

	if err := db.Order("id desc").
		Offset((paging.Page - 1) * paging.Limit).
		Limit(paging.Limit).
		Find(&result).Error; err != nil {
		return nil, common.ErrDB(err)
	}

	return result, nil
}

func (sql *SqlStore) CreateItem(ctx context.Context, data *item.TodoItemCreation) error {
	if err := sql.db.Create(&data).Error; err != nil {
		return common.ErrDB(err)
	}
	return nil
}

func (sql *SqlStore) UpdateItem(ctx context.Context, condition map[string]interface{}, dataUpdate *item.TodoItemUpdate) error {
	if err := sql.db.Where(condition).Updates(dataUpdate).Error; err != nil {
		return common.ErrDB(err)
	}
	return nil
}

func (sql *SqlStore) DeleteItem(ctx context.Context, condition map[string]interface{}) error {
	deletedStatus := item.ItemStatusDeleted

	if err := sql.db.Table(item.TodoItem{}.TableName()).
		Where(condition).
		Updates(map[string]interface{}{
			"status": deletedStatus.String(),
		}).Error; err != nil {
		return common.ErrDB(err)
	}
	return nil
}
