package models

import (
	"time"
	"gorm.io/gorm"
)

type Asset struct {
	ID           uint           `gorm:"primarykey" json:"id"`
	CreatedAt    time.Time      `json:"createdAt"`
	UpdatedAt    time.Time      `json:"updatedAt"`
	DeletedAt    gorm.DeletedAt `gorm:"index" json:"-"`
	Name         string         `json:"name"`
	Category     string         `json:"category"`
	Location     string         `json:"location"`
	Status       string         `json:"status"`
	CurrentStock int            `json:"currentStock"`
	ReorderLevel int            `json:"reorderLevel"`
	Value        float64        `json:"value"`
}
