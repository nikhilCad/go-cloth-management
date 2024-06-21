package controller

import (
	"cloth-management-system/database"
	"cloth-management-system/models"

	"github.com/gin-gonic/gin"

	"go.mongodb.org/mongo-driver/mongo"
)

type OrderItemPack struct {
	Table_id    *string
	Order_items []models.OrderItem
}

var orderItemCollection *mongo.Collection = database.OpenCollection(database.Client, "orderItem")

func GetOrderItems() gin.HandlerFunc {
	return func(c *gin.Context) {
		
	}
}

func GetOrderItemsByOrder() gin.HandlerFunc {
	return func(c *gin.Context) {
		
	}
}


func GetOrderItem() gin.HandlerFunc {
	return func(c *gin.Context) {
		
	}
}

func UpdateOrderItem() gin.HandlerFunc {
	return func(c *gin.Context) {
		
	}
}

func CreateOrderItem() gin.HandlerFunc {
	return func(c *gin.Context) {
		
	}
}