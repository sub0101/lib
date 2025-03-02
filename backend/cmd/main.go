package main

import (
	"fmt"
	"libraryManagement/app/routes"
	"libraryManagement/internal/database"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
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
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}
