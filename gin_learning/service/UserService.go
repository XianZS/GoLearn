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

// PostUser POST User
func PostUser(c *gin.Context) {
	// 传入参数：user对象JSON格式
	// 传入方式：request.body
	user := pojo.User{}
	err := c.BindJSON(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, "Error")
		return
	}
	userList = append(userList, user)
	c.JSON(http.StatusOK, "Successfully posted!")
}

// DeleteUser DELETE User
func DeleteUser(c *gin.Context) {
	// 传入参数：userId
	// 传入方式：URL参数
	// c.Param()获取的是：动态参数在 URL 最后一个阶段
	// 比如说，路由为 /users/some，那么当传入路由为 /users/some/123 时，
	// c.Param("id") 会返回 "123"。
	userId, _ := strconv.Atoi(c.Param("UserId"))
	for i, u := range userList {
		if u.UserId == userId {
			userList = append(userList[:i], userList[i+1:]...)
			c.JSON(http.StatusOK, "Successfully deleted!")
			return
		}
	}
	c.JSON(http.StatusNotFound, "Error")
}

// PutUser PUT User
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
	for i, u := range userList {
		if u.UserId == userId {
			userList[i] = before
			c.JSON(http.StatusOK, "Successfully put!")
			return
		}
	}
	c.JSON(http.StatusNotFound, "Error")
}
