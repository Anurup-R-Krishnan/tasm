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

type ProcurementRequest struct {
	ID                uint           `gorm:"primarykey" json:"id"`
	CreatedAt         time.Time      `json:"createdAt"`
	UpdatedAt         time.Time      `json:"updatedAt"`
	DeletedAt         gorm.DeletedAt `gorm:"index" json:"-"`
	Title             string         `json:"title"`
	Status            string         `json:"status"` // Draft, Pending Approval, PO Raised, Shipping, Received
	Priority          string         `json:"priority"` // Low, Medium, High
	EstimatedValue    float64        `json:"estimatedValue"`
	ActualValue       float64        `json:"actualValue"`
	RequestorName     string         `json:"requestorName"`
	RequestorInitials string         `json:"requestorInitials"`
	Department        string         `json:"department"`
	PONumber          string         `json:"poNumber"`
}

type LedgerEntry struct {
	ID            uint           `gorm:"primarykey" json:"id"`
	CreatedAt     time.Time      `json:"createdAt"`
	UpdatedAt     time.Time      `json:"updatedAt"`
	DeletedAt     gorm.DeletedAt `gorm:"index" json:"-"`
	TransactionID string         `json:"transactionId"`
	Date          time.Time      `json:"date"`
	Description   string         `json:"description"`
	Amount        float64        `json:"amount"`
	Type          string         `json:"type"` // Credit, Debit
	Category      string         `json:"category"`
}

type LeaseAgreement struct {
	ID             uint           `gorm:"primarykey" json:"id"`
	CreatedAt      time.Time      `json:"createdAt"`
	UpdatedAt      time.Time      `json:"updatedAt"`
	DeletedAt      gorm.DeletedAt `gorm:"index" json:"-"`
	LeaseID        string         `json:"leaseId"`
	Vendor         string         `json:"vendor"`
	StartDate      time.Time      `json:"startDate"`
	EndDate        time.Time      `json:"endDate"`
	MonthlyCost    float64        `json:"monthlyCost"`
	Status         string         `json:"status"` // Active, Expired
}

type DepreciationSchedule struct {
	ID             uint           `gorm:"primarykey" json:"id"`
	CreatedAt      time.Time      `json:"createdAt"`
	UpdatedAt      time.Time      `json:"updatedAt"`
	DeletedAt      gorm.DeletedAt `gorm:"index" json:"-"`
	AssetID        string         `json:"assetId"`
	AssetName      string         `json:"assetName"`
	PurchaseValue  float64        `json:"purchaseValue"`
	CurrentValue   float64        `json:"currentValue"`
	Method         string         `json:"method"` // Straight Line, Declining Balance
}

type LedgerEntry struct {
	ID            uint           `gorm:"primarykey" json:"id"`
	CreatedAt     time.Time      `json:"createdAt"`
	UpdatedAt     time.Time      `json:"updatedAt"`
	DeletedAt     gorm.DeletedAt `gorm:"index" json:"-"`
	TransactionID string         `json:"transactionId"`
	Date          time.Time      `json:"date"`
	Description   string         `json:"description"`
	Amount        float64        `json:"amount"`
	Type          string         `json:"type"` // Credit, Debit
	Category      string         `json:"category"`
}

type LeaseAgreement struct {
	ID             uint           `gorm:"primarykey" json:"id"`
	CreatedAt      time.Time      `json:"createdAt"`
	UpdatedAt      time.Time      `json:"updatedAt"`
	DeletedAt      gorm.DeletedAt `gorm:"index" json:"-"`
	LeaseID        string         `json:"leaseId"`
	Vendor         string         `json:"vendor"`
	StartDate      time.Time      `json:"startDate"`
	EndDate        time.Time      `json:"endDate"`
	MonthlyCost    float64        `json:"monthlyCost"`
	Status         string         `json:"status"` // Active, Expired
}

type DepreciationSchedule struct {
	ID             uint           `gorm:"primarykey" json:"id"`
	CreatedAt      time.Time      `json:"createdAt"`
	UpdatedAt      time.Time      `json:"updatedAt"`
	DeletedAt      gorm.DeletedAt `gorm:"index" json:"-"`
	AssetID        string         `json:"assetId"`
	AssetName      string         `json:"assetName"`
	PurchaseValue  float64        `json:"purchaseValue"`
	CurrentValue   float64        `json:"currentValue"`
	Method         string         `json:"method"` // Straight Line, Declining Balance
}
