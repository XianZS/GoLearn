package main

import "fmt"

/*
变量类型解析：

	String:字符串类型
	Integers:整数类型
		int_n:-2**n~2**n-1
		uint_n:0~2**(n+1)-1
		对于int而言，当电脑是32bits时，int长度为32位，同理，当电脑是64bits时，int长度为64位
	Booleans:布尔类型
	Maps:
*/
func main() {
	fmt.Println("variable go things code")
	// 变量属性 变量名 变量类型
	var userName string
	var userId int
	var userNumber float64
	userName = "kom"
	userId = 1001
	userNumber = 1.23
	fmt.Printf("username:%v\n,userIdL:%v\n,number:%v\n", userName, userId, userNumber)
}
