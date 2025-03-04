package utility

import (
	"fmt"
	"log"
	"strconv"

	"github.com/gin-gonic/gin"
)

var RequestTypes = map[string]bool{"issue": true, "return": true}
var IssueStatusTypes = map[string]bool{"accepted": true, "rejected": true}

func GetQueryItem(c *gin.Context, id string) uint {
	idQuery := c.Query(id)
	fmt.Println("idQuery", idQuery)
	retid, _ := strconv.Atoi(idQuery)
	log.Printf("query id %v", retid)
	return uint(retid)
}

func GetContextItem(c *gin.Context, id string) uint {
	userId, _ := c.Get(id)
	result := uint(userId.(float64))
	return result
}
func GetParamItem(c *gin.Context, id string) (uint, bool) {
	idParams, exist := c.Params.Get(id)
	result, _ := strconv.Atoi(idParams)
	if result == 0 {
		return 0, false
	}
	return uint(result), exist

}
