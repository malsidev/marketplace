package repository

import (
	"errors"
	"fmt"
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
		err := tx.Where("user_id = ?", cart.UserID).First(cart).Error

		if errors.Is(err, gorm.ErrRecordNotFound) {
			if err := tx.Create(cart).Error; err != nil {
				return err
			}
		} else if err != nil {
			return err
		}

		item.CartID = cart.ID

		var existing models.CartItem

		err = tx.Where(
			"cart_id = ? AND product_id = ?",
			item.CartID,
			item.ProductID,
		).First(&existing).Error

		if err == nil {
			return tx.Model(&existing).
				Update("quantity", gorm.Expr("quantity + ?", item.Quantity)).Error
		}

		if !errors.Is(err, gorm.ErrRecordNotFound) {
			return err
		}

		return tx.Create(item).Error
	})
}

func (r *CartRepository) GetCart(id int) ([]models.CartItem, error) {
	fmt.Println("Полученный ID:", id)
	var data []models.CartItem

	err := r.db.
		Table("cart_items").
		Select("cart_items.*").
		Joins("JOIN carts ON carts.id = cart_items.cart_id").
		Where("carts.user_id = ?", id).
		Find(&data).Error
	if err != nil {
		return nil, err
	}
	fmt.Printf("Корзина: %+v, ошибка: %v\n", data, err)
	return data, nil
}
