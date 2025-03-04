package requestevent

import (
	"fmt"
	"libraryManagement/internal/dto"
	"libraryManagement/utility"

	utl "libraryManagement/utility"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type RequestEventController struct {
	requestService *RequestEventService
}

var issueStatusType dto.RequestIssueStatus

func NewRequestEventController(db *gorm.DB) *RequestEventController {
	service := RequestEventService{db: db}
	return &RequestEventController{requestService: &service}
}

func (rc *RequestEventController) AddRequest(c *gin.Context) {

	var requestEvent dto.RequestEventDTO
	readerId := utility.GetContextItem(c, "id")
	libId := utility.GetContextItem(c, "libId")

	if err := c.ShouldBindJSON(&requestEvent); err != nil {
		fmt.Println(err)
		utility.SendResponse(c, 400, false, " Bad request", nil, err.Error())
		return
	}
	requestEvent.ReaderID = readerId

	err := rc.requestService.AddRequest(libId, requestEvent)
	if err != nil {
		utility.SendResponse(c, err.Code, false, err.Message, nil, err.Details)
		return

	}
	utility.SendResponse(c, 201, true, "Request Successfully Created", nil)

}

func (rc *RequestEventController) GetAllRequest(c *gin.Context) {

	id := utility.GetContextItem(c, "id")
	libId := utility.GetContextItem(c, "libId")
	response, err := rc.requestService.GetAllRequest(id, libId)

	if err != nil {
		utility.SendResponse(c, err.Code, false, err.Message, nil, err.Details)
		return
	}
	utility.SendResponse(c, 200, true, "successfully fetched all requests", response)
}

func (rc *RequestEventController) GetRequest(c *gin.Context) {

	reqId, valid := utility.GetParamItem(c, "id")
	readerId := utility.GetContextItem(c, "id")
	if !valid {
		utility.SendResponse(c, 400, false, "invalid params", nil, "Params not found")
	}
	respose, err := rc.requestService.GetRequest(readerId, reqId)
	if err != nil {
		utility.SendResponse(c, err.Code, false, err.Message, nil, err.Details)
		return
	}
	utility.SendResponse(c, 200, true, "Successfully Fetched the request", respose)
}

func (rc *RequestEventController) UpdateRequest(c *gin.Context) {

	idParams, valid := c.Params.Get("id")
	userId, _ := c.Get("id")
	libId, _ := c.Get("libId")

	if !valid {
		utl.SendResponse(c, 400, false, "Invalid Params", nil)
		return
	}
	requestId, _ := strconv.Atoi(idParams)

	if err := c.ShouldBindBodyWithJSON(&issueStatusType); err != nil || !utility.IssueStatusTypes[issueStatusType.Type] {
		utl.SendResponse(c, 400, false, "Invalid Body Type", nil, fmt.Errorf("invalid Body").Error())
		return
	}

	err := rc.requestService.UpdateRequest(uint(libId.(float64)), uint(requestId), uint(userId.(float64)), issueStatusType)
	if err != nil {
		utl.SendResponse(c, 400, false, "Unable to Update the request", nil, err.Error())
		return
	}
	utl.SendResponse(c, 200, true, "Successfully Updated the request", nil)

}

func (rc *RequestEventController) GetUserRequests(c *gin.Context) {

	id := utility.GetContextItem(c, "id")

	response, err := rc.requestService.GetUserRequests(id)

	if err != nil {
		utility.SendResponse(c, 400, false, "Not able to fetch requests", nil, err.Error())
		return
	}
	utility.SendResponse(c, 202, true, "successfully fetched all requests", response)

}

func (rc *RequestEventController) DeleteRequest(c *gin.Context) {
	reqId, valid := utility.GetParamItem(c, "id")
	userId := utility.GetContextItem(c, "id")
	if !valid {
		utility.SendResponse(c, 400, false, "No valid params", nil)
		return

	}
	if err := rc.requestService.DeleteRequest(userId, reqId); err != nil {
		utility.SendResponse(c, err.Code, false, err.Message, nil, err.Details)
		return
	}
	utility.SendResponse(c, 200, true, "Succesfully Deleted the Request", nil)
}
