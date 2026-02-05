package domain

import "main/src/features/users/domain/entities"

// AuthRepository defines the persistence operations needed for auth use cases.
type AuthRepository interface {
	Create(user *entities.User) error
	FindByEmail(email string) (*entities.User, error)
}
