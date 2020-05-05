package v1

import (
	res "atlantis/common/response"
	"atlantis/services"

	utils "atlantis/common/utils"

	"github.com/gin-gonic/gin"
)

//AddUserReq 添加用户参数
type AddUserReq struct {
	Username string `json:"Username"`
	Password string `json:"Password"`
}

// AddUser godoc
// @Summary 添加用户
// @Description 添加用户
// @Tags 类别
// @Accept json
// @Produce json
// @Param body body model.LoginReq true  "请求参数"
// @Success 200 {string} string "ok"
// @Router /user/login [post]
func AddUser(c *gin.Context) {
	var req AddUserReq
	user := c.MustGet("remoterUser").(string)
	if err := c.BindJSON(&req); err != nil {
		// 参数解析错误
		res.FailResponse(utils.ParamsErr, err.Error(), c)
	}
	if err := services.UserCreate(req.Username, req.Password, user); err != nil {
		res.SuccessResponse(nil, c)
		return
	}
	res.SuccessResponse(nil, c)

}

// List godoc
// @Summary 列表实例
// @Description 描述信息
// @Tags 类别
// @Accept json
// @Produce json
// @Param limit query string false  "20"
// @Param offset query string false  "0"
// @Success 200 {string} string "ok"
// @Router /ucloud/list [get]
func List(c *gin.Context) {
	res.SuccessResponse(nil, c)
}
