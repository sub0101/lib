package requestevent

import (
	"fmt"
	"libraryManagement/app/middleware"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupRequestEventRouter(router *gin.RouterGroup, db *gorm.DB) {
	requestController := NewRequestEventController(db)
	fmt.Println("inside the request")
	requestEventRouter := router.Group("/request")
	{
		requestEventRouter.Use(middleware.ValidateRefreshToken(db))

		requestEventRouter.POST("/", middleware.IsAuth("reader"), requestController.AddRequest)
		requestEventRouter.GET("/all", middleware.IsAuth("owner", "admin"), requestController.GetAllRequest)
		requestEventRouter.GET("/", middleware.IsAuth("reader"), requestController.GetUserRequests)
		requestEventRouter.GET("/:id", middleware.IsAuth(), requestController.GetRequest)
		requestEventRouter.PATCH("/:id", middleware.IsAuth("owner", "admin"), requestController.UpdateRequest)
		requestEventRouter.DELETE("/:id", middleware.IsAuth("owner", "reader"), requestController.DeleteRequest)
	}

}
