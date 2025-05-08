package pojo

import (
	"GoLearn/gin_learning/database"
	"fmt"
)

type User struct { // Users
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

func CreateUser(user User) User {
	database.DBconnect.Create(&user)
	return user
}

func DeleteUser(userId int) int {
	user := User{}
	result := database.DBconnect.Where("user_id=?", userId).Delete(&user)
	return int(result.RowsAffected)
}

func UpdateUser(userId int, user User) User {
	database.DBconnect.Where("user_id=?", userId).Updates(&user)
	return user
}
