package routes

import (
	"main/src/features/items/infraestructure/controller"

	"github.com/gin-gonic/gin"
)

func RegisterItemRoutes(public *gin.RouterGroup, protected *gin.RouterGroup, ctrl *controller.ItemController) {
	public.GET("/items", ctrl.List)
	public.GET("/items/:id", ctrl.Detail)
	public.GET("/categories", ctrl.Categories)

	protected.POST("/items", ctrl.Create)
	protected.PUT("/items/:id", ctrl.Update)
	protected.DELETE("/items/:id", ctrl.Delete)
	protected.POST("/categories", ctrl.CreateCategory)
	protected.PUT("/categories/:id", ctrl.UpdateCategory)
	protected.DELETE("/categories/:id", ctrl.DeleteCategory)
}
