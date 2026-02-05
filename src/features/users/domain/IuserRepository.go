package domain

import "main/src/features/users/domain/entities"

// UserRepository is the outbound port for user persistence.
type UserRepository interface {
	Create(user *entities.User) error
	FindByEmail(email string) (*entities.User, error)
	FindByID(id uint) (*entities.User, error)
	Update(user *entities.User) error
}
