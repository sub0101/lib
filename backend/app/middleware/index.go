package middleware

import (
	"fmt"
	"libraryManagement/internal/models"
	"libraryManagement/utility"
	"log"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"gorm.io/gorm"
)

func ValidateRefreshToken(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		log.Printf(" refresh middleware")
		authorizationHeader := strings.Split(c.Request.Header.Get("authorization"), " ")
		fmt.Println(authorizationHeader)
		fmt.Println(len(authorizationHeader), "hai")
		if len(authorizationHeader) < 2 || len(authorizationHeader) > 2 {
			utility.SendResponse(c, 401, false, "Bad Request", nil, fmt.Errorf("invalid token").Error())
			c.Abort()
			return
		}

		authorizationHeader[0] = ""
		token := authorizationHeader[1]
		if token == "" {
			utility.SendResponse(c, 401, false, "Unauthorized", nil, fmt.Errorf("invalid token").Error())
			c.Abort()
			return
		}
		jwtToken, err := utility.VerifyToken(token)
		if err != nil {
			log.Printf("error %v", err)
			utility.SendResponse(c, 401, false, "Unauthorized", nil, err.Error())
			c.Abort()
			return
		}
		claims, _ := jwtToken.Claims.(jwt.MapClaims)
		payload := claims["payload"].(map[string]interface{})

		role := payload["Role"]
		log.Printf("user Id %v", payload["Id"])
		result := db.Model(&models.User{}).Where("id=? and role=?", payload["Id"], role).First(&models.User{})
		if result.Error != nil {
			log.Printf("error %v", err)
			utility.SendResponse(c, 401, false, "Unauthorized", nil, fmt.Errorf("Invalid Token").Error())
			c.Abort()
		}
		log.Printf("role %v", role)
		fmt.Printf("valid user")
		// c.Next()
	}
}

func IsAuth(roles ...string) gin.HandlerFunc {

	return func(c *gin.Context) {
		fmt.Println("middlerware")

		authorizationHeader := strings.Split(c.Request.Header.Get("authorization"), " ")

		if len(authorizationHeader) < 2 || len(authorizationHeader) > 2 {
			utility.SendResponse(c, 401, false, "Unauthorized", nil, fmt.Errorf("Invalid Token").Error())
			c.Abort()
			return
		}
		authorizationHeader[0] = ""
		token := authorizationHeader[1]
		jwtToken, err := utility.VerifyToken(token)
		if err != nil {

			utility.SendResponse(c, 401, false, "Unauthorized", nil, err.Error())
			c.Abort()
			return
		}
		claims, _ := jwtToken.Claims.(jwt.MapClaims)
		payload := claims["payload"].(map[string]interface{})

		role := payload["Role"]
		userId := payload["Id"]
		libId := payload["LibId"]

		c.Set("id", userId)
		c.Set("libId", libId)

		isInclude := len(roles) == 0

		for _, r := range roles {

			if r == role {

				isInclude = true
				break
			}
		}

		fmt.Println(isInclude)
		if !isInclude {
			fmt.Println("invalid user")
			utility.SendResponse(c, http.StatusForbidden, false, "Page Not Found", nil)
			c.Abort()
			fmt.Println("invalid user2")
			return
		}

		c.Next()
	}

}
