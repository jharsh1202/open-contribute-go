package routers

import (
	"open-contribute/pkg/api/controllers"
	"open-contribute/pkg/db/repositories"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func RegisterRoutes(router *gin.Engine, db *gorm.DB) {
	// Initialize the repository
	bookRepo := repositories.NewBookRepository(db)

	// Initialize the controller with the repository
	bookController := controllers.NewBookController(bookRepo)

	// Define routes
	routes := router.Group("/books")
	routes.POST("/", bookController.AddBook)
	routes.GET("/", bookController.GetBooks)
	routes.GET("/:id", bookController.GetBook)
	routes.PUT("/:id", bookController.UpdateBook)
	routes.DELETE("/:id", bookController.DeleteBook)
}
