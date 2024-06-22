package common

import "time"

type SqlModel struct {
	Id        int        `json:"id" gorm:"column:id;"`
	CreatedAt *time.Time `json:"createdAt" gorm:"column:created_at;"`
	UpdatedAt *time.Time `json:"updatedAt,omitempty" gorm:"column:updated_at;"`
}
