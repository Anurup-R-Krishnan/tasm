package models

import (
	"time"
	"gorm.io/gorm"
)

type AuditSession struct {
	ID        uint           `gorm:"primarykey" json:"id"`
	CreatedAt time.Time      `json:"createdAt"`
	UpdatedAt time.Time      `json:"updatedAt"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
	Title             string         `json:"title"`
	Status            string         `json:"status"` // e.g. In Progress, Completed
	Auditor           string         `json:"auditor"`
	Location          string         `json:"location"`
	SessionID         string         `json:"sessionId"`
	VerifiedPercent   int            `json:"verifiedPercent"`
	MissingCount      int            `json:"missingCount"`
	UnregisteredCount int            `json:"unregisteredCount"`
}

type Consumable struct {
	ID           uint           `gorm:"primarykey" json:"id"`
	CreatedAt    time.Time      `json:"createdAt"`
	UpdatedAt    time.Time      `json:"updatedAt"`
	DeletedAt    gorm.DeletedAt `gorm:"index" json:"-"`
	Name         string         `json:"name"`
	Category     string         `json:"category"`
	CurrentStock int            `json:"currentStock"`
	ReorderLevel int            `json:"reorderLevel"`
	Location     string         `json:"location"`
}

type MaintenanceContract struct {
	ID          uint           `gorm:"primarykey" json:"id"`
	CreatedAt   time.Time      `json:"createdAt"`
	UpdatedAt   time.Time      `json:"updatedAt"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"-"`
	VendorName  string         `json:"vendorName"`
	ServiceType string         `json:"serviceType"`
	StartDate   time.Time      `json:"startDate"`
	EndDate     time.Time      `json:"endDate"`
	Status      string         `json:"status"` // Active, Expired
}

type WorkOrder struct {
	ID            uint           `gorm:"primarykey" json:"id"`
	CreatedAt     time.Time      `json:"createdAt"`
	UpdatedAt     time.Time      `json:"updatedAt"`
	DeletedAt     gorm.DeletedAt `gorm:"index" json:"-"`
	WorkOrderID   string         `json:"workOrderId"`
	Title         string         `json:"title"`
	AssetLocation string         `json:"assetLocation"`
	Issue         string         `json:"issue"`
	Severity      string         `json:"severity"` // Critical, High, Medium, Low
	TargetDate    time.Time      `json:"targetDate"`
	Status        string         `json:"status"` // Open, In Progress, Closed
}
