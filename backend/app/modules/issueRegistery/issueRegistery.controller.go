package issueregistery

import (
	"fmt"
	"libraryManagement/utility"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type IssueRegistryController struct {
	issueService *IssueRegistryService
}

func NewIssueRegistryController(db *gorm.DB) *IssueRegistryController {

	service := IssueRegistryService{DB: db}
	return &IssueRegistryController{issueService: &service}
}

func (ic *IssueRegistryController) GetAllIssue(c *gin.Context) {

	libId := utility.GetContextItem(c, "libId")
	result, err := ic.issueService.GetAllIssueRequest(libId)
	if err != nil {
		utility.SendResponse(c, err.Code, false, err.Message, nil, err.Details)
		return
	}
	fmt.Println("all issued fetched")
	utility.SendResponse(c, 202, true, "Successfully fetched all issue", result)

}
