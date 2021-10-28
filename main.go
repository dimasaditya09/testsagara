package main

import (
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"log"
	"os"
	"testSagara/config"
	"testSagara/database"
	"testSagara/models"
	"testSagara/repositories"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	//DATABASE CONNECTION
	db, err := database.Connect()
	if err != nil {
		log.Fatalln(err)
	}

	//MIGRATION
	db.AutoMigrate(
		&models.Product{},
		&models.User{},
	)

	productRepository := repositories.NewProductRepository(db)

	//ROUTE
	route := gin.Default()
	config.AuthRoutes(route)
	config.ProductRoutes(productRepository, route)

	// NO ROUTE
	route.NoRoute(func(c *gin.Context) {
		var d struct{}
		c.JSON(404, gin.H{"message": "Page not found", "success": false, "data": d})
	})
	route.Run(":" + os.Getenv("PORT"))
}
