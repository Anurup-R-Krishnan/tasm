package models

import (
	"gorm.io/gorm"
	"time"
)

type SystemConfig struct {
	ID        uint           `gorm:"primarykey" json:"id"`
	CreatedAt time.Time      `json:"createdAt"`
	UpdatedAt time.Time      `json:"updatedAt"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
	Key       string         `gorm:"uniqueIndex" json:"key"`
	Value     string         `json:"value"`
}
