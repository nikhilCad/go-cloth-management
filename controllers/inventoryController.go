package controller

import (

	"github.com/gin-gonic/gin"
)

var inventoryCollection *mongo.Collection = database.OpenCollection(database.Client, "inventory")

func GetInventories() gin.HandlerFunc {
	return func(c *gin.Context) {
		
	}
}

func GetInventory() gin.HandlerFunc {
	return func(c *gin.Context) {
		
	}
}

func CreateInventory() gin.HandlerFunc {
	return func(c *gin.Context) {
		
	}
}

func UpdateInventory() gin.HandlerFunc {
	return func(c *gin.Context) {
		
	}
}