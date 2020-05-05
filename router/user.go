package apirouter

import (
	v1 "atlantis/api/v1"

	"github.com/gin-gonic/gin"
)

//InitUserRouter  用户路由
func InitUserRouter(group *gin.RouterGroup) {
	userRouter := group.Group("user")
	userRouter.POST("login", v1.AddUser)
}
