package database

import (
	"log"
	"time"

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

	var db *gorm.DB
	var lastErr error
	for i := 0; i < 10; i++ {
		db, lastErr = gorm.Open(postgres.Open(config.DSN()), &gorm.Config{})
		if lastErr == nil {
			break
		}
		log.Printf("Failed to connect to database (attempt %d/10): %v. Retrying in 2s...", i+1, lastErr)
		time.Sleep(2 * time.Second)
	}

	if lastErr != nil {
		return lastErr
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
		&models.SystemConfig{},
	); err != nil {
		log.Fatalf("Failed to auto-migrate models: %v", err)
	}

	return nil
}
