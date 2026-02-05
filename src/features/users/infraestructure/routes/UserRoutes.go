package routes

import (
	UserController "main/src/features/users/infraestructure/controller"

	"github.com/gin-gonic/gin"
)

func RegisterUserRoutes(rg *gin.RouterGroup, ctrl *UserController.UserController) {
	api := rg.Group("/users")
	api.GET("/me", ctrl.GetMe)
	api.PUT("/me", ctrl.UpdateMe)
}
