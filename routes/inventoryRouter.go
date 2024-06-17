package routes

import (
	controller "cloth-management-system/controllers"

	"github.com/gin-gonic/gin"
)

func InventoryRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.GET("/inventories", controller.GetInventories())
	incomingRoutes.GET("/inventories/:inventory_id", controller.GetInventory())
	incomingRoutes.POST("/inventories", controller.CreateInventory())
	incomingRoutes.PATCH("/inventories/:inventory_id", controller.UpdateInventory())
}