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
	Name           string         `json:"name"`
	TagID          string         `json:"tagId"`
	Category       string         `json:"category"`
	Location       string         `json:"location"`
	Status         string         `json:"status"`
	Custodian      string         `json:"custodian"`
	CurrentStock   int            `json:"currentStock"`
	ReorderLevel   int            `json:"reorderLevel"`
	Value          float64        `json:"value"`
	PurchaseDate   time.Time      `json:"purchaseDate"`
	WarrantyStatus string         `json:"warrantyStatus"`
	WarrantyExpiry time.Time      `json:"warrantyExpiry"`
}
