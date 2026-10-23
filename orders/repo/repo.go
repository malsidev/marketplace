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

func (r *OrderRepo) PostOrder(order *models.Orders, items *models.OrdersItems) error {
	return r.db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(order).Error; err != nil {
			return err
		}

		items.OrderID = order.ID

		if err := tx.Create(items).Error; err != nil {
			return err
		}

		return nil
	})
}
