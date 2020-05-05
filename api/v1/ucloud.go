package v1

import (
	res "go-project-init/common/response"

	"github.com/gin-gonic/gin"
)

//ImageList 镜像列表实例
func ImageList(c *gin.Context) {
	res.SuccessResponse(nil, c)
}
