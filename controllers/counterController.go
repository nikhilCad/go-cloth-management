package controller

import (

	"github.com/gin-gonic/gin"
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