package domain

import "main/src/features/users/domain/entities"

type AuthRepository interface {
	Create(user *entities.User) error
	FindByEmail(email string) (*entities.User, error)
}
