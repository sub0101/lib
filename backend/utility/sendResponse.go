package utility

import (
	"github.com/gin-gonic/gin"
)

func SendResponse(c *gin.Context, statusCode int, success bool, message string, data interface{}, err ...string) {

	response := gin.H{
		"success": success,
		"message": message,
	}
	if len(err) > 0 {
		response["error"] = err[0]
	}
	if data != nil {
		response["data"] = data
	}

	c.JSON(statusCode, response)

}
