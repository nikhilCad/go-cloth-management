package routes

import (
	controller "cloth-management-system/controllers"

	"github.com/gin-gonic/gin"
)

func CounterRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.GET("/counters", controller.GetCounters())
	incomingRoutes.GET("/counters/:counter_id", controller.GetCounter())
	incomingRoutes.POST("/counters", controller.CreateCounter())
	incomingRoutes.PATCH("/counters/:counter_id", controller.UpdateCounter())
}