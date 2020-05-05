package apirouter

import (
	v1 "atlantis/api/v1"

	"github.com/gin-gonic/gin"
)

//InitURouter  用户路由
func InitURouter(group *gin.RouterGroup) {
	userRouter := group.Group("ucloud")
	userRouter.GET("list", v1.List)
}
