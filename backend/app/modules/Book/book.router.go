package book

import (
	"fmt"
	"libraryManagement/app/middleware"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupBookRouter(router *gin.RouterGroup, db *gorm.DB) {
	bookController := NewBookController(db)
	fmt.Println("book router intialized")
	bookRouter := router.Group("/books")
	{
		bookRouter.Use(middleware.ValidateRefreshToken(db))

		bookRouter.GET("/", middleware.IsAuth(), bookController.GetAllBook)
		bookRouter.GET("/issued", middleware.IsAuth("reader"), bookController.GetIssuedBooks)
		bookRouter.GET("/:id", middleware.IsAuth(), bookController.GetBook)
		bookRouter.POST("/", middleware.IsAuth("owner", "admin"), bookController.AddBook)
		bookRouter.PATCH("/:id", middleware.IsAuth("owner", "admin"), bookController.UpdateBook)
		bookRouter.DELETE("/:id", middleware.IsAuth("owner", "admin"), bookController.DeleteBook)
		bookRouter.GET("/search", middleware.IsAuth(), bookController.SearchBooks)
	}
}
