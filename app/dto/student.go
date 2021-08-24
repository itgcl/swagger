package dto

import "swagger/model"

type RequestStudentCreate struct {
	Username    string `json:"username" binding:"required" example:"张三"`
	Password    string `json:"password" binding:"required" example:"123456"`
	Age         int    `json:"age" binding:"required" example:"10"`
	Gender      string `json:"gender" binding:"required" example:"男"`
	Class       int    `json:"class" binding:"required" example:"5"`
	HomeAddress string `json:"home_address" example:"上海市徐汇区xxx路"`
}

type ReplyStudentCreate struct {
	StudentID model.IDStudent `json:"studentID" example:"1"`
}

type RequestStudentUpdate struct {
	Username    string `json:"username" binding:"required" example:"李四"`
	Password    string `json:"password" binding:"required" example:"654321"`
	Age         int    `json:"age" binding:"required" example:"12"`
	Gender      string `json:"gender" binding:"required" example:"女"`
	Class       int    `json:"class" binding:"required" example:"5"`
	HomeAddress string `json:"home_address" example:"上海市闵行区xxx路"`
}

type ReplyStudentInfo struct {
	Username    string `json:"username" binding:"required" example:"李四"`
	Password    string `json:"password" binding:"required" example:"654321"`
	Age         int    `json:"age" binding:"required" example:"12"`
	Gender      string `json:"gender" binding:"required" example:"女"`
	Class       int    `json:"class" binding:"required" example:"5"`
	HomeAddress string `json:"home_address" example:"上海市闵行区xxx路"`
}
