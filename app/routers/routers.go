package routers

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"swagger/app/middleware"
	v1 "swagger/app/routers/v1"
)

func InitRouters() *gin.Engine {
	r := gin.New()

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	{
		r.POST("/student/api/v1/login", v1.Login)
		apiv1 := r.Group("/student/api/v1", middleware.ValidateAuth())

		// 学生系统CRUD
		{
			apiv1.GET("/:studentID", v1.Info)
			apiv1.POST("/create", v1.Create)
			apiv1.PATCH("/:studentID", v1.Update)
			apiv1.DELETE("/:studentID", v1.Delete)
		}
	}
	return r
}