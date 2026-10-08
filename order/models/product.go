package models

type Orders struct {
	ID       int `gorm:"primarykey"`
	Quantity int `gorm:"not null"`
}
