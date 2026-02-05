package entities

import "time"

type User struct {
	ID            uint
	FullName      string
	Email         string
	Password      string
	Phone         string
	Address       string
	ProfilePicURL string
	Role          string
	CreatedAt     time.Time
	UpdatedAt     time.Time
}

func NewUser(fullName, email, password, phone, address, profilePicURL, role string) *User {
	if role == "" {
		role = "user"
	}
	return &User{
		FullName:      fullName,
		Email:         email,
		Password:      password,
		Phone:         phone,
		Address:       address,
		ProfilePicURL: profilePicURL,
		Role:          role,
		CreatedAt:     time.Now(),
		UpdatedAt:     time.Now(),
	}
}
