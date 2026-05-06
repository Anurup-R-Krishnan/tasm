package models

import (
	"gorm.io/gorm"
	"time"
)

// AssetEvent tracks all changes to assets for history/audit purposes
type AssetEvent struct {
	ID                uint           `gorm:"primarykey" json:"id"`
	CreatedAt         time.Time      `json:"createdAt"`
	UpdatedAt         time.Time      `json:"updatedAt"`
	DeletedAt         gorm.DeletedAt `gorm:"index" json:"-"`
	AssetID           uint           `json:"assetId"`
	EventType         string         `json:"eventType"` // CHECKOUT, CHECKIN, STATUS_CHANGE, CUSTODIAN_CHANGE, LOCATION_CHANGE
	Description       string         `json:"description"`
	PreviousStatus    string         `json:"previousStatus"`
	NewStatus         string         `json:"newStatus"`
	PreviousCustodian string         `json:"previousCustodian"`
	NewCustodian      string         `json:"newCustodian"`
	ActorID           uint           `json:"actorId"`
	ActorName         string         `json:"actorName"`
	DueDate           *time.Time     `json:"dueDate,omitempty"`
	Notes             string         `json:"notes"`
}
