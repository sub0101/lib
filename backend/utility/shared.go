package utility

import (
	"strconv"

	"github.com/gin-gonic/gin"
)

var RequestTypes = map[string]bool{"issue": true, "return": true}
var IssueStatusTypes = map[string]bool{"accepted": true, "rejected": true}

func GetQueryItem(c *gin.Context, id string) uint {
	idQuery := c.Query(id)
	libId, _ := strconv.Atoi(idQuery)
	return uint(libId)
}

func GetContextItem(c *gin.Context, id string) uint {
	userId, _ := c.Get(id)
	result := uint(userId.(float64))
	return result
}
func GetParamItem(c *gin.Context, id string) (uint, bool) {
	idParams, exist := c.Params.Get(id)
	result, _ := strconv.Atoi(idParams)
	return uint(result), exist

}
