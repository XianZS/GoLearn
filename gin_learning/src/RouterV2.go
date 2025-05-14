package src

import (
	"GoLearn/gin_learning/middlewares"
	"GoLearn/gin_learning/service"
	"fmt"
	"github.com/gin-gonic/gin"
)

func AddUserRouterV2(r *gin.RouterGroup) {
	// 创建总路由组
	yyy := r.Group("/read", middlewares.SetSession())
	fmt.Println("read")
	yyy.POST("/login", service.LoginUser)
	yyy.Use(middlewares.AuthSession())
	{
		yyy.GET("/read", service.Read)
	}
}
