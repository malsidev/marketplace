package db

import (
	"fmt"
	"log"
	"order/config"
	"order/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewPostgres(cfg *config.Config) (*gorm.DB, error) {
	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		cfg.DBHost,
		cfg.DBPort,
		cfg.DBUser,
		cfg.DBPassword,
		cfg.DBName,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("не удалось подлючиться к бд: ", err)
		return nil, err
	}

	_ = db.Migrator().DropTable(&models.Orders{}, &models.OrdersItems{})
	if err := db.AutoMigrate(&models.Orders{}, &models.OrdersItems{}); err != nil {
		return nil, err
	}

	return db, nil
}
