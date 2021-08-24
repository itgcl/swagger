package v1

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"swagger/app/dto"
)

// @Summary 登录
// @Tags 用户
// @version 1.0
// @Accept json
// @Produce json
// @Param {object} body dto.RequestLogin true "请求参数"
// @Success 200 {object} dto.ReplyLogin "请求成功"
// @Failure 400 "参数错误"
// @Failure 401 "没有权限"
// @Failure 500 "服务器错误"
// @Router /login [post]
func Login(c *gin.Context)  {
	request := new(dto.RequestLogin)
	if err := c.BindJSON(request); err != nil {
		c.Status(http.StatusBadRequest)
		return
	}
	reply := dto.ReplyLogin{Token: "123456"}
	c.JSON(http.StatusOK, reply)
	return
}