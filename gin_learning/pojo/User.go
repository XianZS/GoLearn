package pojo

import (
	"GoLearn/gin_learning/database"
	"fmt"
)

type User struct {
	UserId       int    `json:"userId"`
	UserName     string `json:"userName"`
	UserPassWord string `json:"userPassword"`
	UserEmail    string `json:"userEmail"`
}

func FindAllUsers() []User {
	var users []User
	database.DBconnect.Find(&users)
	fmt.Println("=========")
	fmt.Println(users)
	fmt.Println("=========")
	return users
}