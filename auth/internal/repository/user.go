package repository

type User struct {
	Id           int
	Phone        string
	PasswordHash string
}

type UserRepository interface {
	Create(user User) error
	GetByUsername(phone string) (*User, error)
}
