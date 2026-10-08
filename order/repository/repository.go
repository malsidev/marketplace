package repository

import (
	"order/models"

	"gorm.io/gorm"
)

type Order struct {
	ID       int
	Quantity string
}

type OrderRepository struct {
	db *gorm.DB
}

func NewOrderRepository(db *gorm.DB) *OrderRepository {
	return &OrderRepository{
		db: db,
	}
}

func (r *OrderRepository) Create(data *models.Orders) error {
	return r.db.Create(data).Error
}
