package database

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"tasm-backend/models"
)

var DB *gorm.DB

func ConnectDB() error {
	config, err := LoadConfig()
	if err != nil {
		return err
	}

	db, err := gorm.Open(postgres.Open(config.DSN()), &gorm.Config{})
	if err != nil {
		return err
	}

	DB = db
	log.Printf("Database connected: %s:%s/%s", config.Host, config.Port, config.Name)

	if err := DB.AutoMigrate(
		&models.Asset{},
		&models.AuditSession{},
		&models.AuditDiscrepancy{},
		&models.Consumable{},
		&models.MaintenanceContract{},
		&models.WorkOrder{},
		&models.ProcurementRequest{},
		&models.LedgerEntry{},
		&models.LeaseAgreement{},
		&models.DepreciationSchedule{},
		&models.SoftwareLicense{},
		&models.SystemUser{},
		&models.UserRole{},
		&models.Reservation{},
		&models.Location{},
		&models.SystemAlert{},
	); err != nil {
		log.Fatalf("Failed to auto-migrate models: %v", err)
	}

	return nil
}
