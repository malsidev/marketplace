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

func (r *CartRepository) Create(data *models.Carts) error {
	return r.db.Create(data).Error
}

func (r *CartRepository) CreateCartWithItem(cart *models.Carts, item *models.CartItem) error {
	return r.db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(cart).Error; err != nil {
			return err
		}

		item.CartID = cart.ID

		if err := tx.Create(item).Error; err != nil {
			return err
		}

		return nil
	})
}
