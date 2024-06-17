package controller

import (
	"github.com/gin-gonic/gin"
)

var clothCollection *mongo.Collection = database.OpenCollection(database.Client, "cloth")
var validate = validator.New()

func GetCloths() gin.HandlerFunc {
	return func(c *gin.Context) {

		
	}
}

func GetCloth() gin.HandlerFunc {
	return func(c *gin.Context) {
		
	}
}

func CreateCloth() gin.HandlerFunc {
	return func(c *gin.Context) {
		
	}
}

func UpdateCloth() gin.HandlerFunc {
	return func(c *gin.Context) {
		
	}
}