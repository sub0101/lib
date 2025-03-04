package main

// @title Library Management API
// @version 1.0
// @description API for managing a library.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8000
// @BasePath /api/v1
// @schemes http

import (
	"fmt"
	"libraryManagement/app/routes"
	"libraryManagement/internal/database"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	_ "libraryManagement/app/modules/auth"
	_ "libraryManagement/docs"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func main() {

	server := gin.Default()
	godotenv.Load(".env")
	err, db := database.InitDB()
	if err != nil {
		server.Use(func(ctx *gin.Context) {
			ctx.AbortWithStatus(501)
		})
	}
	server.Use(CORSMiddleware())

	// server.Use(ValidateRefreshToken(db))
	// server.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	server.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	router := server.Group("/api/v1")
	routes.SetupRouter(router, db)

	server.Run(":8000")

}

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("inside cors")
		c.Writer.Header().Set("Access-Control-Allow-Origin", "http://localhost:5173")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT,PATCH, DELETE")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}

// https://documenter.getpostman.com/view/38592065/2sAYdkGUCa
