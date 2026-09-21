package service

import (
	"time"
	
	"github.com/golang-jwt/jwt/v5"
)

type JWTService struct {
	secret []byte
}

type Claims struct {
	UserID int `json:"sub"`
	Type string `json:"type"`
	jwt.RegisteredClaims
}

func NewJWTService(secret string) *JWTService {
	return &JWTService{
		secret: []byte(secret),
	}
}

func (j *JWTService) GenerateToken(userID int) (string, error) {
	claims := Claims{
		UserID: userID,
		Type:   "access",
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(15 * time.Minute)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			Issuer:    "marketplace-auth",
		},
	}
	token := jwt.NewWithClaims(
		jwt.SigningMethodHS256,
		claims,
	)

	return token.SignedString(j.secret)
}