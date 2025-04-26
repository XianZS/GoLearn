package main

import "fmt"

func main() {
	fmt.Println("conference")
	// 普通变量定义为 var，可以被修改
	var name = "jom"
	name = "kom"
	// const 常量定义为 const，不可以被修改
	const number = 10
	// 占位符使用,使用fmt.Printf("%v",value)
	fmt.Printf("number:%v", number)
	// 打印输出
	fmt.Println(number)
	fmt.Println(name)
}
