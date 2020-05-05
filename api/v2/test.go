package v2

import (
	res "go-project-init/common/response"

	"github.com/gin-gonic/gin"
)

//TxImageList  腾讯镜像列表
func TxImageList(c *gin.Context) {
	res.SuccessResponse(nil, c)
}
