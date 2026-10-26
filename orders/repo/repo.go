package repo

import (
	"fmt"
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

func (r *OrderRepo) PostOrder(order *models.Orders, items []models.OrdersItems) error {
	return r.db.Transaction(func(tx *gorm.DB) error {
		totalprice := 0.0

		for i := range items {
			totalprice += (items[i].Price * float64(items[i].Quantity))
		}

		order.TotalPrice = totalprice
		if err := tx.Create(order).Error; err != nil {
			return err
		}

		for i := range items {
			items[i].OrderID = order.ID
			if err := tx.Create(&items[i]).Error; err != nil {
				return err
			}
		}
		return nil
	})
}

func (r *OrderRepo) GetProduct(id int) ([]models.Orders, error) {
	var orders []models.Orders

	err := r.db.
		Select("id", "user_id", "status", "total_price", "created_at").
		Preload("OrderItem").
		Where("user_id = ?", id).
		Find(&orders).Error

	if err != nil {
		return nil, err
	}
	fmt.Println(orders)

	return orders, nil
}
