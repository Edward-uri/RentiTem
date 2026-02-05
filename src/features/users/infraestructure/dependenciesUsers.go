package infraestructure

import (
	"main/src/features/users/application"
	"main/src/features/users/domain"
	"main/src/features/users/infraestructure/controller"
	"main/src/features/users/infraestructure/persistence"

	"gorm.io/gorm"
)

type UsersDependencies struct {
	Repo         domain.UserRepository
	CreateUserUC *application.CreateUserUseCase
	GetProfileUC *application.GetProfileUseCase
	UpdateUC     *application.UpdateProfileUseCase
	Controller   *controller.UserController
}

func NewUsersDependencies(db *gorm.DB) UsersDependencies {
	repo := persistence.NewGormUserRepository(db)

	return UsersDependencies{
		Repo:         repo,
		CreateUserUC: application.NewCreateUserUseCase(repo),
		GetProfileUC: application.NewGetProfileUseCase(repo),
		UpdateUC:     application.NewUpdateProfileUseCase(repo),
		Controller:   controller.NewUserController(application.NewGetProfileUseCase(repo), application.NewUpdateProfileUseCase(repo)),
	}
}
