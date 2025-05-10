package src

import (
	"GoLearn/gin_learning/service"
	"fmt"
	"github.com/gin-gonic/gin"
)

func AddUserRouterV2(r *gin.RouterGroup) {
	// 创建总路由组
	yyy := r.Group("/read")
	fmt.Println("read")
	yyy.GET("/read", service.Read)
}
