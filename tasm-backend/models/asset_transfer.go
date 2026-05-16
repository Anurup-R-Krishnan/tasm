package models

import "time"

// AssetTransfer records every asset transfer for chain of custody
type AssetTransfer struct {
	ID               uint       `gorm:"primarykey" json:"id"`
	CreatedAt        time.Time  `json:"createdAt"`
	AssetID          uint       `json:"assetId"`
	FromEntity       string     `json:"fromEntity"`
	ToEntity         string     `json:"toEntity"`
	FromLocation     string     `json:"fromLocation"`
	ToLocation       string     `json:"toLocation"`
	Reason           string     `json:"reason"`
	EffectiveDate    time.Time  `json:"effectiveDate"`
	ApprovedBy       string     `json:"approvedBy"`
	ApprovalStatus   string     `json:"approvalStatus"` // Pending, Approved, Rejected
	RequiresApproval bool       `json:"requiresApproval"`
	Reverted         bool       `json:"reverted"`
	RevertedAt       *time.Time `json:"revertedAt"`
	RevertedBy       string     `json:"revertedBy"`
	ActorID          uint       `json:"actorId"`
	ActorName        string     `json:"actorName"`
}

// AssetDisposalRecord captures the full disposal details when an asset is disposed
type AssetDisposalRecord struct {
	ID                      uint      `gorm:"primarykey" json:"id"`
	CreatedAt               time.Time `json:"createdAt"`
	AssetID                 uint      `json:"assetId"`
	DisposalMethod          string    `json:"disposalMethod"` // Sell, Scrap, Recycle, Donate
	ResidualValue           float64   `json:"residualValue"`
	DecommissionDate        time.Time `json:"decommissionDate"`
	ValuationDate           time.Time `json:"valuationDate"`
	EnvironmentalCompliance bool      `json:"environmentalCompliance"`
	DataWipingConfirmed     bool      `json:"dataWipingConfirmed"`
	ComplianceVerified      bool      `json:"complianceVerified"`
	Notes                   string    `json:"notes"`
	ActorID                 uint      `json:"actorId"`
	ActorName               string    `json:"actorName"`
}
