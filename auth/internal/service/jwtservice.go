package service

import (
	"time"
	
	"github.com/golang-jwt/jwt/v5"
)

type JWTService struct {
	secret []byte
}

func NewJWTService(secret string) *JWTService {
	return &JWTService{
		secret: [byte(secret)]
	}
}

func (j *JWTService) GenerateToken(userID uint) (string, error) {
	claims := jwt.MapClaims{
		"sub" : userID,
		"iat": time.Now().Unix(),
		"exp": time.Now().Add(15 * time.Minute).Unix(),
	}
	token := jwt.NewWithClaims(
		jwt.SigningMethodHS256,
		claims
	)

	return token.SignedString(j.secret)
}