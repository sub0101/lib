package routes

import (
	"fmt"

	book "libraryManagement/app/modules/Book"
	library "libraryManagement/app/modules/Library"
	user "libraryManagement/app/modules/User"
	"libraryManagement/app/modules/auth"
	issueregistery "libraryManagement/app/modules/issueRegistery"
	requestevent "libraryManagement/app/modules/requestEvent"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupRouter(router *gin.RouterGroup, db *gorm.DB) {
	fmt.Println("setup the main router")
	auth.SetupAuthRouter(router, db)

	library.SetupLibraryRouter(router, db)
	book.SetupBookRouter(router, db)
	requestevent.SetupRequestEventRouter(router, db)
	user.SetupUserRouter(router, db)
	issueregistery.SetupIssueRegistryRouter(router, db)

}
