package src

import (
	"GoLearn/gin_learning/service"
	"github.com/gin-gonic/gin"
)

func AddUserRouter(r *gin.RouterGroup) {
	user := r.Group("/users")
	user.GET("/FindAllUser", service.FindAllUser)
	user.POST("/PostUser", service.PostUser)
	user.DELETE("/DeleteUser/:UserId", service.DeleteUser)
	user.PUT("/PutUser/:UserId", service.PutUser)
}
