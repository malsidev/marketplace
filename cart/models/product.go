package models

import "time"

type Carts struct {
	ID        int `gorm:"primarykey"`
	UserID    int
	CreatedAt time.Time
	UpdatedAt time.Time

	CartItems []CartItem `gorm:"foreignKey:CartID"`
}

type CartItem struct {
	ID        int `gorm:"primaryKey"`
	CartID    int
	ProductID int
	Quantity  int
}
