package controller

import (
	"context"
	"cloth-management-system/database"
	"cloth-management-system/models"
	"time"
	"net/http"
	"log"
	"fmt"

	"github.com/gin-gonic/gin"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
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

		// same  as cloth controller

		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		inventoryId := c.Param("inventory_id")
		var inventory models.Inventory

		err := inventoryCollection.FindOne(ctx, bson.M{"inventory_id": inventoryId}).Decode(&inventory)
		defer cancel()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "error occured while fetching the inventory"})
		}
		c.JSON(http.StatusOK, inventory)
		
	}
}

func CreateInventory() gin.HandlerFunc {
	return func(c *gin.Context) {

		// same as cloth controller

		var inventory models.Inventory
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)

		if err := c.BindJSON(&inventory); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		validationErr := validate.Struct(inventory)
		if validationErr != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": validationErr.Error()})
			return
		}

		inventory.Created_at, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
		inventory.Updated_at, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
		inventory.ID = primitive.NewObjectID()
		inventory.Inventory_id = inventory.ID.Hex()

		result, insertErr := inventoryCollection.InsertOne(ctx, inventory)
		if insertErr != nil {
			msg := fmt.Sprintf("Inventory item was not created")
			c.JSON(http.StatusInternalServerError, gin.H{"error": msg})
			return
		}
		defer cancel()
		c.JSON(http.StatusOK, result)
		defer cancel()
		
	}
}

func inTimeSpan(start, end, check time.Time) bool {
	return start.After(time.Now()) && end.After(start)
}

func UpdateInventory() gin.HandlerFunc {
	return func(c *gin.Context) {

		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		var inventory models.Inventory

		if err := c.BindJSON(&inventory); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		inventoryId := c.Param("inventory_id")
		filter := bson.M{"inventory_id": inventoryId}

		var updateObj primitive.D

		if inventory.Start_Date != nil && inventory.End_Date != nil {
			if !inTimeSpan(*inventory.Start_Date, *inventory.End_Date, time.Now()) {
				msg := "kindly retype the time"
				c.JSON(http.StatusInternalServerError, gin.H{"error": msg})
				defer cancel()
				return
			}

			updateObj = append(updateObj, bson.E{"start_date", inventory.Start_Date})
			updateObj = append(updateObj, bson.E{"end_date", inventory.End_Date})

			if inventory.Name != "" {
				updateObj = append(updateObj, bson.E{"name", inventory.Name})
			}
			if inventory.Category != "" {
				updateObj = append(updateObj, bson.E{"name", inventory.Category})
			}

			inventory.Updated_at, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
			updateObj = append(updateObj, bson.E{"updated_at", inventory.Updated_at})

			upsert := true

			opt := options.UpdateOptions{
				Upsert: &upsert,
			}

			result, err := inventoryCollection.UpdateOne(
				ctx,
				filter,
				bson.D{
					{"$set", updateObj},
				},
				&opt,
			)

			if err != nil {
				msg := "Inventory update failed"
				c.JSON(http.StatusInternalServerError, gin.H{"error": msg})
			}

			defer cancel()
			c.JSON(http.StatusOK, result)
		}
		
	}
}