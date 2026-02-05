package main

import (
	"log"

	"github.com/gin-gonic/gin"

	"main/src/core/config"
	coredb "main/src/core/db"
	AuthInfra "main/src/features/auth/infraestructure"
	AuthRoutes "main/src/features/auth/infraestructure/routes"
	UserPersistence "main/src/features/users/infraestructure/persistence"
)

func main() {
	cfg := config.Load()

	database, err := coredb.New(cfg.DBURL)
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}

	// Auto-migrate user table.
	if err := coredb.AutoMigrate(database, &UserPersistence.UserModel{}); err != nil {
		log.Fatalf("failed to migrate database: %v", err)
	}

	r := gin.Default()
	api := r.Group("/api/v1")

	authDeps := AuthInfra.NewDependencies(database)
	AuthRoutes.RegisterAuthRoutes(api, authDeps.Controller)

	if err := r.Run(":" + cfg.Port); err != nil {
		log.Fatalf("server error: %v", err)
	}
}
