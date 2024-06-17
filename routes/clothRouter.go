package routes

import (
	controller "cloth-management-system/controllers"

	"github.com/gin-gonic/gin"
)

func ClothRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.GET("/cloths", controller.GetCloths())
	incomingRoutes.GET("/cloths/:cloth_id", controller.GetCloth())
	incomingRoutes.POST("/cloths", controller.CreateCloth())
	incomingRoutes.PATCH("/cloths/:cloth_id", controller.UpdateCloth())
}