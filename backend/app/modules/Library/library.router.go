package library

import (
	"libraryManagement/app/middleware"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupLibraryRouter(router *gin.RouterGroup, db *gorm.DB) {

	libraryController := NewLibraryController(db)
	libraryRouter := router.Group("/library")
	{
		libraryRouter.Use(middleware.ValidateRefreshToken(db))

		libraryRouter.GET("/", middleware.IsAuth(), libraryController.GetAllLibrary)
		libraryRouter.POST("/addLibrary", middleware.IsAuth(), libraryController.AddLibrary)
		libraryRouter.GET("/:id", middleware.IsAuth(), libraryController.GetLibrary)
		libraryRouter.PATCH("/:id", middleware.IsAuth("owner"), libraryController.UpdateLibrary)
		libraryRouter.DELETE("/:id", middleware.IsAuth("owner"), libraryController.DeleteLibrary)
	}

}
