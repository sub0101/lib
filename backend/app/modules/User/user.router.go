package user

import (
	"libraryManagement/app/middleware"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupUserRouter(router *gin.RouterGroup, db *gorm.DB) {
	userController := NewUserController(db)
	userRouter := router.Group("/users")
	{
		userRouter.Use(middleware.ValidateRefreshToken(db))

		userRouter.PATCH("/", userController.UpdateUser)
		userRouter.GET("/:id", userController.GetUser)
		userRouter.GET("/", middleware.IsAuth("owner", "admin"), userController.GetAllUser)
		userRouter.PATCH("/make_admin", middleware.IsAuth("owner"), userController.MakeAdmin)
	}
}
