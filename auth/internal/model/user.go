package model

type User struct {
	Id int `gorm:"primarykey"`
	Phone string `gorm:"unique;not null"`
	PasswordHash string `gorm:"not null"`
}