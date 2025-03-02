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
		if len(authorizationHeader) <= 1 && len(authorizationHeader) < 2 {
			utility.SendResponse(c, 400, false, "Bad Request", nil, fmt.Errorf("Invalid Token").Error())
			c.Abort()
		}

		authorizationHeader[0] = ""

		token := authorizationHeader[1]
		jwtToken, err := utility.VerifyToken(token)
		if err != nil {
			utility.SendResponse(c, 401, false, "Unauthorized", nil, err.Error())
			c.Abort()
		}
		claims, _ := jwtToken.Claims.(jwt.MapClaims)
		payload := claims["payload"].(map[string]interface{})

		role := payload["Role"]
		result := db.Model(&models.User{}).Where("id=? and role=?", payload["Id"], role).First(&models.User{})
		if result.Error != nil {
			utility.SendResponse(c, 401, false, "Unauthorized", nil, fmt.Errorf("Invalid Token").Error())
			c.Abort()
		}
		log.Printf("role %v", role)
		c.Next()
	}
}

func IsAuth(roles ...string) gin.HandlerFunc {

	return func(c *gin.Context) {
		fmt.Println("middlerware")
		// fmt.Println(c.Request.Header.Get("authorization"))
		authorizationHeader := strings.Split(c.Request.Header.Get("authorization"), " ")
		if len(authorizationHeader) <= 1 && len(authorizationHeader) < 2 {
			utility.SendResponse(c, 400, false, "Bad Request", nil, fmt.Errorf("Invalid Token").Error())
			c.Abort()
		}
		fmt.Println(authorizationHeader)
		authorizationHeader[0] = ""
		fmt.Println(len(authorizationHeader))
		token := authorizationHeader[1]
		jwtToken, err := utility.VerifyToken(token)
		if err != nil {
			utility.SendResponse(c, 401, false, "Unauthorized", nil, err.Error())
			c.Abort()
		}
		claims, _ := jwtToken.Claims.(jwt.MapClaims)
		payload := claims["payload"].(map[string]interface{})

		role := payload["Role"]
		userId := payload["Id"]
		libId := payload["LibId"]
		log.Printf("role %v", role)
		fmt.Println(userId)
		c.Set("id", userId)
		c.Set("libId", libId)

		isInclude := len(roles) == 0

		for _, r := range roles {

			if r == role {

				isInclude = true
				break
			}
		}

		if !isInclude {
			utility.SendResponse(c, http.StatusForbidden, false, "Page Not Found", nil)
			c.Abort()
		}

	}

}
