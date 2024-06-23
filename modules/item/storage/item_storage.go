package storage

import (
	"context"
	"go-demo/common"
	"go-demo/modules/item/model"
)

func (sql *SqlStore) ListItem(ctx context.Context, filter *model.Filter, paging *common.Paging, moreKeys ...string) ([]model.TodoItem, error) {
	var result []model.TodoItem

	db := sql.db.Where("status <> ?", "Deleted")

	if f := filter; f != nil {
		if v := f.Status; v != "" {
			db = db.Where("status = ?", v)
		}
	}

	if err := db.Table(model.TodoItem{}.TableName()).Count(&paging.Total).Error; err != nil {
		return nil, err
	}

	if err := db.Order("id desc").
		Offset((paging.Page - 1) * paging.Limit).
		Limit(paging.Limit).
		Find(&result).Error; err != nil {
		return nil, err
	}

	return result, nil
}

func (sql *SqlStore) GetItem(ctx context.Context, condition map[string]interface{}) (*model.TodoItem, error) {
	var data model.TodoItem
	if err := sql.db.Where(condition).First(&data).Error; err != nil {
		return nil, err
	}
	return &data, nil
}

func (sql *SqlStore) CreateItem(ctx context.Context, data *model.TodoItemCreation) error {
	if err := sql.db.Create(&data).Error; err != nil {
		return err
	}
	return nil
}

func (sql *SqlStore) UpdateItem(ctx context.Context, condition map[string]interface{}, dataUpdate *model.TodoItemUpdate) error {
	if err := sql.db.Where(condition).Updates(dataUpdate).Error; err != nil {
		return err
	}
	return nil
}

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
