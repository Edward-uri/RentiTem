package services

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

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

func (s *JWTService) Generate(userID uint, email, role string) (string, error) {
	claims := jwt.MapClaims{
		"sub":   userID,
		"email": email,
		"role":  role,
		"iss":   s.issuer,
		"exp":   time.Now().Add(s.expiresIn).Unix(),
		"iat":   time.Now().Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(s.secret))
}

func (s *JWTService) Validate(tokenStr string) (uint, string, string, error) {
	token, err := jwt.Parse(tokenStr, func(t *jwt.Token) (interface{}, error) {
		if t.Method != jwt.SigningMethodHS256 {
			return nil, errors.New("unexpected signing method")
		}
		return []byte(s.secret), nil
	})
	if err != nil {
		return 0, "", "", err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return 0, "", "", errors.New("invalid token")
	}

	sub, ok := claims["sub"].(float64)
	if !ok {
		return 0, "", "", errors.New("invalid sub")
	}
	email, ok := claims["email"].(string)
	if !ok {
		return 0, "", "", errors.New("invalid email")
	}
	role, ok := claims["role"].(string)
	if !ok {
		return 0, "", "", errors.New("invalid role")
	}

	return uint(sub), email, role, nil
}
