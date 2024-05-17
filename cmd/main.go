package main

import (
	"fmt"
	"log"
	"os"

	router "open-contribute/pkg/api/routers"
	"open-contribute/pkg/db"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func main() {
	// Load environment variables from .env file
	viper.SetConfigFile(".env")
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("Error reading config file: %v", err)
	}

	// Fetch environment variables
	port := viper.GetString("PORT")
	dbPath := viper.GetString("DB_PATH")

	// Create a new Gin engine with default middleware
	gin_router := gin.Default()

	// Initialize Database
	dbHandler, err := db.Init(dbPath)
	if err != nil {
		log.Fatalf("Error initializing database: %v", err)
	}
	// defer dbHandler.Close()

	// Register Routes
	router.RegisterRoutes(gin_router, dbHandler)

	// Default route to display port and database path
	gin_router.GET("/", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"port":   port,
			"dbPath": dbPath,
		})
	})

	// Start the server
	log.Printf("Starting server on port %s", port)
	err = gin_router.Run(fmt.Sprintf(":%s", port))
	if err != nil {
		log.Fatalf("Error starting server: %v", err)
		os.Exit(1)
	}
}
