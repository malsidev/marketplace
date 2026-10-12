package models

type Cart struct {
	ID       int `gorm:"primarykey"`
	Quantity int `gorm:"not null"`
}
