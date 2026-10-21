package repo

import (
	"order/models"

	"gorm.io/gorm"
)

type OrderRepo struct {
	db *gorm.DB
}

func NewOrderRepo(db *gorm.DB) *OrderRepo {
	return &OrderRepo{db: db}
}

func (r *OrderRepo) GetOrder() ([]models.OrdersItems, error) {
	var items []models.OrdersItems
	err := r.db.Find(&items).Error
	return items, err
}
