package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

//Result 返回结构体
type Result struct {
	Success bool
	Code    int
	Msg     string
	Data    interface{}
}

func response(success bool, code int, msg string, data interface{}, c *gin.Context) {
	r := Result{success, code, msg, data}
	c.JSON(http.StatusOK, r)
}

//SuccessResponse 成功返回
func SuccessResponse(data interface{}, c *gin.Context) {
	response(true, 0, "success", data, c)
}

//FailResponse  失败返回
func FailResponse(code int, msg string, c *gin.Context) {
	response(false, code, msg, nil, c)
}
