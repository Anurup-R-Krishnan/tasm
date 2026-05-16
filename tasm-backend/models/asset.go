package models

import (
	"gorm.io/gorm"
	"time"
)

type Asset struct {
	ID            uint           `gorm:"primarykey" json:"id"`
	CreatedAt     time.Time      `json:"createdAt"`
	UpdatedAt     time.Time      `json:"updatedAt"`
	DeletedAt     gorm.DeletedAt `gorm:"index" json:"-"`
	Name          string         `json:"name"`
	TagID         string         `json:"tagId"`
	SerialNumber  string         `json:"serialNumber"`
	Department    string         `json:"department"`
	Category      string         `json:"category"`
	Location      string         `json:"location"`
	Condition     string         `json:"condition"`
	Status        string         `json:"status"`
	Custodian     string         `json:"custodian"`
	CurrentStock  int            `json:"currentStock"`
	ReorderLevel  int            `json:"reorderLevel"`
	Value         float64        `json:"value"`
	PurchaseDate  time.Time      `json:"purchaseDate"`
	PurchasePrice float64        `json:"purchasePrice"`

	// Warranty
	WarrantyStatus    string    `json:"warrantyStatus"`
	WarrantyExpiry    time.Time `json:"warrantyExpiry"`
	WarrantyStartDate time.Time `json:"warrantyStartDate"`
	WarrantyEndDate   time.Time `json:"warrantyEndDate"`
	WarrantyProvider  string    `json:"warrantyProvider"`
	WarrantyTerms     string    `json:"warrantyTerms"`

	// Receipt
	ReceiptFilePath string `json:"receiptFilePath"`

	// Depreciation
	UsefulLifeYears    int     `json:"usefulLifeYears"`
	DepreciationMethod string  `json:"depreciationMethod"` // Straight Line, Declining Balance
	ResidualValue      float64 `json:"residualValue"`
	ReplacementCost    float64 `json:"replacementCost"`

	// Lifecycle
	LifecycleStatus string `json:"lifecycleStatus"` // Procurement, Deployed, Under Maintenance, End of Life, Disposed

	// Disposal
	DisposalMethod          string     `json:"disposalMethod"` // Sell, Scrap, Recycle, Donate
	DisposalDate            *time.Time `json:"disposalDate"`
	DisposalNotes           string     `json:"disposalNotes"`
	EnvironmentalCompliance bool       `json:"environmentalCompliance"`
	DataWipingConfirmed     bool       `json:"dataWipingConfirmed"`
	// Financial Links
	LeaseID uint `json:"leaseId,omitempty"`
}
