package auth

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupAuthRouter(router *gin.RouterGroup, db *gorm.DB) {

	fmt.Println("auth router setup")
	authController := NewAuthController(db)
	authRouter := router.Group("/auth")
	{
		authRouter.POST("/login", authController.Login)
		authRouter.POST("/signup", authController.Signup)
		authRouter.POST("/library/signup", authController.SignupLibrary)
		authRouter.POST("/forget", func(ctx *gin.Context) {})
	}

}
