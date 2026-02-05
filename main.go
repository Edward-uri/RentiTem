package main

import (
	"log"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"main/src/core/config"
	coredb "main/src/core/db"
	_ "main/src/docs"
	AuthInfra "main/src/features/auth/infraestructure"
	AuthMiddleware "main/src/features/auth/infraestructure/middleware"
	AuthRoutes "main/src/features/auth/infraestructure/routes"
	ItemInfra "main/src/features/items/infraestructure"
	ItemPersistence "main/src/features/items/infraestructure/persistence"
	ItemRoutes "main/src/features/items/infraestructure/routes"
	UserInfra "main/src/features/users/infraestructure"
	UserPersistence "main/src/features/users/infraestructure/persistence"
	UserRoutes "main/src/features/users/infraestructure/routes"
)

func main() {
	cfg := config.Load()

	database, err := coredb.New(cfg.DBURL)
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}
	if err := coredb.AutoMigrate(database, &UserPersistence.UserModel{}, &ItemPersistence.CategoryModel{}, &ItemPersistence.ItemModel{}); err != nil {
		log.Fatalf("failed to migrate database: %v", err)
	}

	r := gin.Default()
	api := r.Group("/api/v1")

	authDeps := AuthInfra.NewDependencies(database)
	AuthRoutes.RegisterAuthRoutes(api, authDeps.Controller)

	usersDeps := UserInfra.NewUsersDependencies(database)
	protected := api.Group("")
	protected.Use(AuthMiddleware.JWTAuthMiddleware(authDeps.JWT))

	itemDeps := ItemInfra.NewDependencies(database, usersDeps.Repo, cfg.UploadDir)
	ItemRoutes.RegisterItemRoutes(api, protected, itemDeps.Controller)
	UserRoutes.RegisterUserRoutes(protected, usersDeps.Controller)

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	if err := r.Run(":" + cfg.Port); err != nil {
		log.Fatalf("server error: %v", err)
	}
}
