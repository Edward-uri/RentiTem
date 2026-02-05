package infraestructure

import (
	"os"
	"time"

	"main/src/features/auth/application"
	"main/src/features/auth/infraestructure/controller"
	"main/src/features/auth/infraestructure/services"
	"main/src/features/users/infraestructure/persistence"

	"gorm.io/gorm"
)

// Dependencies bundles all auth adapters and use cases.
type Dependencies struct {
	Controller *controller.AuthController
}

// NewDependencies builds auth feature wiring using shared DB.
func NewDependencies(db *gorm.DB) Dependencies {
	pwd := services.NewBcryptService(12)
	jwtSecret := os.Getenv("JWT_SECRET")
	if jwtSecret == "" {
		jwtSecret = "change-me"
	}
	jwtSvc := services.NewJWTService(jwtSecret, 72*time.Hour, "rentitems")

	repo := persistence.NewGormUserRepository(db)
	uc := application.NewAuthUseCase(repo, pwd, jwtSvc)
	ctrl := controller.NewAuthController(uc)

	return Dependencies{Controller: ctrl}
}
