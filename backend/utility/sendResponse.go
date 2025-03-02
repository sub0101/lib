package utility

import (
	"encoding/json"
	"libraryManagement/internal/dto"
	"libraryManagement/internal/models"

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

func ConvertToDTO(library models.Library) (dto.ResponseGetLibrary, error) {
	var response dto.ResponseGetLibrary

	// Convert `models.Library` → JSON → `dto.ResponseGetLibrary`
	jsonData, err := json.Marshal(library)
	if err != nil {
		return response, err
	}
	err = json.Unmarshal(jsonData, &response)
	return response, err
}
