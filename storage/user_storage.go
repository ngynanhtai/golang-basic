package storage

import (
	"context"
	"errors"
	"go-demo/common"
	"go-demo/models"
	"gorm.io/gorm"
)

func (sql *SqlStore) CreateUser(ctx context.Context, data *models.UserCreation) error {
	if err := sql.db.Create(&data).Error; err != nil {
		return common.ErrDB(err)
	}
	return nil
}

func (sql *SqlStore) GetUser(ctx context.Context, condition map[string]interface{}) (*models.User, error) {
	var data models.User
	if err := sql.db.Where(condition).First(&data).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, common.RecordNotFound
		}
		return nil, common.ErrDB(err)
	}
	return &data, nil
}
