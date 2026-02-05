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
	CreatedAt     time.Time
	UpdatedAt     time.Time
}

func NewUser(fullName, email, password, phone, address, profilePicURL string) *User {
	return &User{
		FullName:      fullName,
		Email:         email,
		Password:      password,
		Phone:         phone,
		Address:       address,
		ProfilePicURL: profilePicURL,
		CreatedAt:     time.Now(),
		UpdatedAt:     time.Now(),
	}
}
