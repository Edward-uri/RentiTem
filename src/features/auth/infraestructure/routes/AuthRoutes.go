package routes

import (
	"main/src/features/auth/infraestructure/controller"

	"github.com/gin-gonic/gin"
)

// RegisterAuthRoutes binds auth endpoints to the router group.
func RegisterAuthRoutes(rg *gin.RouterGroup, ctrl *controller.AuthController) {
	api := rg.Group("/auth")
	api.POST("/register", ctrl.Register)
	api.POST("/login", ctrl.Login)
}
