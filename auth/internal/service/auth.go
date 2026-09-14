package sevice


import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
	"unicode"
)

type AuthService struct {

}

func NewAuthService() *AuthService{
	return &AuthService{}
}

func (s *AuthService) Register(phone string, password string) error {
	if len(phone) != 11 {
		return fmt.Errorf("phone number must be 11 digits")
	}

	if len(password) < 8 {
		return fmt.Errorf("password must be at least 8 characters long")
	}

	hasDigit := false

	for _, r := range password {
		if unicode.IsDigit(r) {
			hasDigit = true
		}
	}

	if hasDigit == false {
		return fmt.Errorf("There must be at least one digit.")
	}

		_, err := bcrypt.GenerateFromPassword(
		[]byte(password),
		bcrypt.DefaultCost,
	)

	if err != nil {
		return fmt.Errorf("failed to hash password")
	}

	return nil
}