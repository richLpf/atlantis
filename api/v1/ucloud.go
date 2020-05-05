package v1

import (
	res "atlantis/common/response"

	"github.com/gin-gonic/gin"
)

//ImageList 镜像列表实例
func ImageList(c *gin.Context) {
	res.SuccessResponse(nil, c)
}
