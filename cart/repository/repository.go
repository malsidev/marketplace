package repository

import (
	"order/models"

	"gorm.io/gorm"
)

type Cart struct {
	ID       int
	Quantity string
}

type CartRepository struct {
	db *gorm.DB
}

func NewCartRepository(db *gorm.DB) *CartRepository {
	return &CartRepository{
		db: db,
	}
}

func (r *CartRepository) Create(data *models.Cart) error {
	return r.db.Create(data).Error
}
