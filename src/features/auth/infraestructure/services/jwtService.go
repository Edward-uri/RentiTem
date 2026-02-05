package services

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

// JWTService generates JWT tokens with a secret and expiration.
type JWTService struct {
	secret    string
	expiresIn time.Duration
	issuer    string
}

func NewJWTService(secret string, expiresIn time.Duration, issuer string) *JWTService {
	if expiresIn == 0 {
		expiresIn = 72 * time.Hour
	}
	return &JWTService{secret: secret, expiresIn: expiresIn, issuer: issuer}
}

func (s *JWTService) Generate(userID uint, email string) (string, error) {
	claims := jwt.MapClaims{
		"sub":   userID,
		"email": email,
		"iss":   s.issuer,
		"exp":   time.Now().Add(s.expiresIn).Unix(),
		"iat":   time.Now().Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(s.secret))
}
