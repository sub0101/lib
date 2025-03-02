package issueregistery

import (
	"libraryManagement/app/middleware"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupIssueRegistryRouter(router *gin.RouterGroup, db *gorm.DB) {

	issueController := NewIssueRegistryController(db)

	IssueRouter := router.Group("/issue_register")
	{
		IssueRouter.Use(middleware.ValidateRefreshToken(db))

		IssueRouter.GET("/", middleware.IsAuth(), issueController.GetAllIssue)
		IssueRouter.GET("/:id", middleware.IsAuth(), issueController.GetIssue)
	}

}
