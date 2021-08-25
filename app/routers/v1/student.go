package v1

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"swagger/app/dto"
)

// @Summary 创建
// @Tags 学生管理
// @Version 1.0
// @Produce json
// @Security ApiKeyAuth
// @Param {object} body dto.RequestStudentCreate true "请求参数"
// @Success 200 {object} dto.ReplyStudentCreate "请求成功"
// @Failure 400 "参数错误"
// @Failure 401 "没有权限"
// @Failure 500 "服务器错误"
// @Router /create [post]
func Create(c *gin.Context)  {
	request := new(dto.RequestStudentCreate)
	if err := c.BindJSON(request); err != nil {
		c.Status(http.StatusBadRequest)
		return
	}

	// 此处省略service操作
	// ......
	// 此处省略service操作

	// TODO 模拟返回
	reply := dto.ReplyStudentCreate{StudentID: 2}
	c.JSON(http.StatusOK, reply)
	return
}

// @Summary 编辑
// @Tags 学生管理
// @version 1.0
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param studentID path int true "学生id"
// @Param {object} body dto.RequestStudentUpdate true "请求参数"
// @Success 204 "请求成功"
// @Failure 400 "参数错误"
// @Failure 401 "没有权限"
// @Failure 404 "访问不存在的信息"
// @Failure 500 "服务器错误"
// @Router /{studentID} [patch]
func Update(c *gin.Context)  {
	request := new(dto.RequestStudentUpdate)
	if err := c.BindJSON(request); err != nil {
		c.Status(http.StatusBadRequest)
		return
	}
	studentID, err := strconv.Atoi(c.Param("studentID"))
	if err != nil {
		c.Status(http.StatusBadRequest)
		return
	}
	// 此处省略service操作
	_ = studentID
	// ......
	// 此处省略service操作

	c.Status(http.StatusNoContent)
	return
}

// @Summary 详情
// @Tags 学生管理
// @Version 1.0
// @Produce json
// @Security ApiKeyAuth
// @Param studentID path int true "学生id"
// @Success 200 {object} dto.ReplyStudentInfo "请求成功"
// @Failure 400 "参数错误"
// @Failure 401 "没有权限"
// @Failure 404 "访问不存在的信息"
// @Failure 500 "服务器错误"
// @Router /{studentID} [get]
func Info(c *gin.Context) {
	studentID, err := strconv.Atoi(c.Param("studentID"))
	if err != nil {
		c.Status(http.StatusBadRequest)
		return
	}

	// 此处省略service操作
	_ = studentID
	// ......
	// 此处省略service操作

	reply := dto.ReplyStudentInfo{
		Username:    "王五",
		Password:    "password",
		Age:         10,
		Gender:      "男",
		Class:       3,
		HomeAddress: "上海市闵行区xx路",
	}
	c.JSON(http.StatusOK, reply)
	return
}

// @Summary 删除
// @Tags 学生管理
// @Version 1.0
// @Security ApiKeyAuth
// @Param studentID path int true "学生id"
// @Success 204 "请求成功"
// @Failure 400 "参数错误"
// @Failure 401 "没有权限"
// @Failure 500 "服务器错误"
// @Router /{studentID} [delete]
func Delete (c *gin.Context)  {
	studentID, err := strconv.Atoi(c.Param("studentID"))
	if err != nil {
		c.Status(http.StatusBadRequest)
		return
	}

	// 此处省略service操作
	_ = studentID
	// ......
	// 此处省略service操作

	c.Status(http.StatusNoContent)
	return
}