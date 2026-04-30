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

	sqlDB, err := db.DB()
	if err != nil {
		return err
	}

	if err := sqlDB.Ping(); err != nil {
		return err
	}

	DB = db
	log.Printf("Database connection established for %s@%s:%s/%s", config.User, config.Host, config.Port, config.Name)

	if err := DB.AutoMigrate(
		&models.Asset{},
		&models.AuditSession{},
		&models.Consumable{},
		&models.MaintenanceContract{},
		&models.WorkOrder{},
	); err != nil {
		return err
	}

	return nil
}
