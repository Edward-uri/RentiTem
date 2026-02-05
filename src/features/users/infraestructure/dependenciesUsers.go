package infraestructure

import (
	"main/src/features/users/application"
	"main/src/features/users/domain"
	"main/src/features/users/infraestructure/persistence"

	"gorm.io/gorm"
)

// UsersDependencies wires the user feature adapters and use cases.
type UsersDependencies struct {
	Repo         domain.UserRepository
	CreateUserUC *application.CreateUserUseCase
}

// NewUsersDependencies builds the concrete implementations for the user feature.
func NewUsersDependencies(db *gorm.DB) UsersDependencies {
	repo := persistence.NewGormUserRepository(db)

	return UsersDependencies{
		Repo:         repo,
		CreateUserUC: application.NewCreateUserUseCase(repo),
	}
}
