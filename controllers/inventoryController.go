package controller

import (
	"context"
	"cloth-management-system/database"
	"time"
	"net/http"
	"log"

	"github.com/gin-gonic/gin"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/bson"
)

var inventoryCollection *mongo.Collection = database.OpenCollection(database.Client, "inventory")

func GetInventories() gin.HandlerFunc {
	return func(c *gin.Context) {

		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)

		// Not findone but find, finds all
		result, err := inventoryCollection.Find(context.TODO(), bson.M{})
		defer cancel()

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "error occured while listing the inventory items"})
		}


		var allInventories []bson.M
		if err = result.All(ctx, &allInventories); err != nil {
			log.Fatal(err)
		}

		// return allInventories in context with status OK
		c.JSON(http.StatusOK, allInventories)
		
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