package src

import (
	"GoLearn/gin_learning/service"
	"github.com/gin-gonic/gin"
)

func AddUserRouterV1(r *gin.RouterGroup) {
	// 创建总路由组
	user := r.Group("/users")
	// 得到所有用户信息
	user.GET("/FindAllUser", service.FindAllUser)
	// 增加用户信息
	user.POST("/PostUser", service.PostUser)
	// 删除用户信息
	user.DELETE("/DeleteUser/:UserId", service.DeleteUser)
	// 修改用户信息
	user.PUT("/PutUser/:UserId", service.PutUser)
}
