package config

import (
	"tcl_be/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {

	dsn := "host=localhost user=postgres password=postgres dbname=inventory_db port=5433 sslmode=disable"

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("Failed to connect database")
	}

	DB = db

	DB.AutoMigrate(
		&models.Product{},
		&models.Inventory{},
		&models.StockIn{},
		&models.StockInItem{},
		&models.StockOut{},
		&models.StockOutItem{},
		&models.StockLog{},
	)
}
