package controller

import (

	"cloth-management-system/database"

	"github.com/gin-gonic/gin"

	"go.mongodb.org/mongo-driver/mongo"
)

var counterCollection *mongo.Collection = database.OpenCollection(database.Client, "counter")

func GetCounters() gin.HandlerFunc {
	return func(c *gin.Context) {
		
	}
}

func GetCounter() gin.HandlerFunc {
	return func(c *gin.Context) {
		
	}
}

func CreateCounter() gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}

func UpdateCounter() gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}