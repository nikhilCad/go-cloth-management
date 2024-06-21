package controller

import (
	"context"
	"time"
	"cloth-management-system/database"
	"cloth-management-system/models"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"

	"go.mongodb.org/mongo-driver/mongo"
	
)

// openCollection fucntion from our own datatbse package
var clothCollection *mongo.Collection = database.OpenCollection(database.Client, "cloth")
var validate = validator.New()

func GetCloths() gin.HandlerFunc {
	return func(c *gin.Context) {

		
	}
}

func GetCloth() gin.HandlerFunc {
	return func(c *gin.Context) {
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		// get cloth id from the context
		clothId := c.Param("cloth_id")
		var cloth models.Cloth

		// find one object with this cloth_id and decode that into our cloth variable we declared above
		err := clothCollection.FindOne(ctx, bson.M{"cloth_id": clothId}).Decode(&cloth)
		defer cancel()

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "error occured while fetching the cloth item"})
		}
		c.JSON(http.StatusOK, cloth)
		
	}
}

func CreateCloth() gin.HandlerFunc {
	return func(c *gin.Context) {

		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		var inventory models.Inventory
		var cloth models.Cloth

		if err := c.BindJSON(&cloth); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		validationErr := validate.Struct(cloth)
		if validationErr != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": validationErr.Error()})
			return
		}
		
		// does this inventory id exist?
		err := inventoryCollection.FindOne(ctx, bson.M{"inventory_id": cloth.Inventory_id}).Decode(&inventory)
		defer cancel()
		if err != nil {
			msg := fmt.Sprintf("inventory was not found")
			c.JSON(http.StatusInternalServerError, gin.H{"error": msg})
			return
		}


		// RFC 3339 is a time format
		cloth.Created_at, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
		cloth.Updated_at, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
		cloth.ID = primitive.NewObjectID()
		cloth.Cloth_id = cloth.ID.Hex()
		var num = toFixed(*cloth.Price, 2)
		cloth.Price = &num

		// insert this above data into our collection
		result, insertErr := clothCollection.InsertOne(ctx, cloth)
			
		if insertErr != nil {
			msg := fmt.Sprintf("Cloth item was not created")
			c.JSON(http.StatusInternalServerError, gin.H{"error": msg})
			return
		}
		defer cancel()
		c.JSON(http.StatusOK, result)
		
	}
}

func UpdateCloth() gin.HandlerFunc {
	return func(c *gin.Context) {
		
	}
}