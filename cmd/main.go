package main

import (
	router "open-contribute/pkg/api/routers"
	"open-contribute/pkg/db"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func main() {
	// SET ENVIRONMENT CONFIG FILE PATH
	viper.SetConfigFile(".env")

	// READ ENVIRONMENT CONFIG FILE
	viper.ReadInConfig()

	// FETCH RELATED ENVIRONMENT VARIABLES
	port := viper.Get("PORT").(string)
	dbPath := viper.Get("DB_PATH").(string)

	// creates a new Gin engine with the default middleware attached which includes Logger and Recovery middleware
	// holds the instance of the Gin engine used to define routes and start the server.
	gin_router := gin.Default()

	// Initialize Database
	dbHandler := db.Init(dbPath)

	// Register Routes
	router.RegisterRoutes(gin_router, dbHandler)

	gin_router.GET("/", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"port": port,
			// "dbUrl": dbUrl,
			"dbPath": dbPath,
		})
	})

	gin_router.Run(port)
}
