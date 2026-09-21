package service

import (
	"errors"
	"fmt"
	"gorm.io/gorm"

	"unicode"
	"auth/internal/model"
	"golang.org/x/crypto/bcrypt"
	"auth/internal/repository"
)

type AuthService struct {
	userRepository *repository.UserRepository
	jwtService     *JWTService
}

func NewAuthService(
	userRepository *repository.UserRepository,
	jwtService *JWTService,
) *AuthService {
	return &AuthService{
		userRepository: userRepository,
		jwtService:     jwtService,
	}
}

func (s *AuthService) Register(phone string, password string) error {
	user, err := s.userRepository.GetByPhone(phone)

	if err == nil {
		return fmt.Errorf("invalid phone or password")
	}

	if !errors.Is(err, gorm.ErrRecordNotFound) {
		// Это уже настоящая ошибка базы данных
		return fmt.Errorf("failed to check phone: %w", err)
	}

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

	passwordHash, err := bcrypt.GenerateFromPassword(
		[]byte(password),
		bcrypt.DefaultCost,
	)

	if err != nil {
		return fmt.Errorf("failed to hash password")
	}

	user = &model.User{
		Phone:        phone,
		PasswordHash: string(passwordHash),
	}

	return s.userRepository.Create(user)
}

func (s *AuthService) Login(phone string, password string) (string, error) {


	if len(phone) != 11 {
		return "", fmt.Errorf("phone number must be 11 digits")
	}

	if len(password) < 8 {
		return "", fmt.Errorf("password must be at least 8 characters long")
	}

	hasDigit := false

	for _, r := range password {
		if unicode.IsDigit(r) {
			hasDigit = true
		}
	}

	if hasDigit == false {
		return "", fmt.Errorf("There must be at least one digit.")
	}


	user, err := s.userRepository.GetByPhone(phone)

	if err != nil {
		return "", fmt.Errorf("the phone number is busy")
	}

	err = bcrypt.CompareHashAndPassword(
		[]byte(user.PasswordHash),
		[]byte(password),
	)

	if err != nil {
		return "", fmt.Errorf("failed to hash password")
	}

	token, err := s.jwtService.GenerateToken(user.Id)
	if err != nil {
		return "", fmt.Errorf("failed to generate token")
	}

	return token, nil
}
