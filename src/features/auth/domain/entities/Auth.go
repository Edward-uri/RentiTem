package entities

import "time"

type Auth struct {
	Token     string
	UserID    uint
	Email     string
	IssuedAt  time.Time
	ExpiresAt time.Time
}
