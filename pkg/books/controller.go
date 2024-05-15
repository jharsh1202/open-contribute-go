package books

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// The handler struct serves as a container for holding a GORM database connection
type handler struct {
	DB *gorm.DB // pointer to a GORM database connection
}

func RegisterRoutes(router *gin.Engine, db *gorm.DB) {
	// Create a handler instance with the database connection
	h := &handler{
		DB: db,
	}

	routes := router.Group("/books")
	routes.POST("/", h.AddBook)
	routes.GET("/", h.GetBooks)
	routes.GET("/:id", h.GetBook)
	routes.PUT("/:id", h.UpdateBook)
	routes.DELETE("/:id", h.DeleteBook)
}
