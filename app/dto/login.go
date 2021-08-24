package dto

type RequestLogin struct {
	Username string `form:"username" binding:"required" example:"zhangsan"`
	Password string `form:"password" binding:"required" example:"123456"`
}

type ReplyLogin struct {
	Token string `json:"token" example:"ethgdjkfgbvdig34jh5g4jhbvhj"`
}
