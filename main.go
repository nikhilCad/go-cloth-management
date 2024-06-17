package main

import (
	"os"

	"cloth-management-system/database"

	middleware "cloth-management-system/middleware"
	routes "cloth-management-system/routes"

	"github.com/gin-gonic/gin"

	"go.mongodb.org/mongo-driver/mongo"
)

// reads if present, creates if not present
var clothCollection *mongo.Collection = database.OpenCollection(database.Client, "cloth")

func main() {
	port := os.Getenv("PORT")

	if port == "" {
		// default port
		port = "8000"
	}

	// Using gin router for JWT 
	router := gin.New()
	router.Use(gin.Logger())
	routes.UserRoutes(router)
	router.Use(middleware.Authentication())

	// Initializing routes
	// functions defined in routes folder
	routes.ClothRoutes(router)
	routes.InventoryRoutes(router)
	routes.CounterRoutes(router)
	routes.OrderRoutes(router)
	routes.OrderItemRoutes(router)
	routes.InvoiceRoutes(router)

	router.Run(":" + port)
}