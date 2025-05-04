package service

import (
	"GoLearn/gin_learning/pojo"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

var userList []pojo.User

// FindAllUser GET User
func FindAllUser(c *gin.Context) {
	// 传入参数: nil
	// 传入方式：nil
	users := pojo.FindAllUsers()
	fmt.Println("*********")
	fmt.Println(users)
	c.JSON(http.StatusOK, users)
}

// PostUser POST User:增加用户
func PostUser(c *gin.Context) {
	// 传入参数：user对象JSON格式
	// 传入方式：request.body
	user := pojo.User{}
	err := c.BindJSON(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, "Error")
		return
	}
	newUser := pojo.CreateUser(user)
	c.JSON(http.StatusOK, newUser)
}

// DeleteUser DELETE User:删除用户
func DeleteUser(c *gin.Context) {
	// 传入参数：userId
	// 传入方式：URL参数
	// c.Param()获取的是：动态参数在 URL 最后一个阶段
	// 比如说，路由为 /users/some，那么当传入路由为 /users/some/123 时，
	// c.Param("id") 会返回 "123"。
	userId, _ := strconv.Atoi(c.Param("UserId"))
	delUser := pojo.DeleteUser(userId)
	if delUser.UserId == 0 {
		c.JSON(http.StatusNotFound, "Not Found")
		return
	} else {
		c.JSON(http.StatusOK, "Successfully delete")
	}
}

// PutUser PUT User:修改用户
func PutUser(c *gin.Context) {
	// 传入参数：userId, user对象JSON格式
	// 传入方式：URL参数, request.body
	// 创建修改对象
	before := pojo.User{}
	// 绑定修改对象
	err := c.BindJSON(&before)
	if err != nil {
		c.JSON(http.StatusBadRequest, "Error")
		return
	}
	userId, _ := strconv.Atoi(c.Param("UserId"))
	newUser := pojo.UpdateUser(userId, before)
	c.JSON(http.StatusNotFound, newUser)
}
