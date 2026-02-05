package routes

import (
	"main/src/features/auth/infraestructure/controller"

	"github.com/gin-gonic/gin"
)

func RegisterAuthRoutes(rg *gin.RouterGroup, ctrl *controller.AuthController) {
	api := rg.Group("/auth")
	api.POST("/register", ctrl.Register)
	api.POST("/login", ctrl.Login)
}
