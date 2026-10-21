package models

import (
	"time"
)

type Orders struct {
	ID         int `gorm:"primarykey"`
	UserID     int
	Status     string
	TotalPrice int
	CreatedAt  time.Time

	OrderItem []OrdersItems `gorm:"foreignKey:OrderID"`
}

type OrdersItems struct {
	ID        int `gorm:"primarykey"`
	OrderID   int
	ProductID int
	Quantity  int
	Price     int
}

/*
CREATE TABLE carts (
    id BIGSERIAL PRIMARY KEY,
    user_id BIGINT NOT NULL UNIQUE,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE TABLE cart_items (
    id BIGSERIAL PRIMARY KEY,
    cart_id BIGINT NOT NULL REFERENCES carts(id) ON DELETE CASCADE,
    product_id BIGINT NOT NULL,
    quantity INTEGER NOT NULL CHECK (quantity > 0),
    UNIQUE (cart_id, product_id)
);
*/
