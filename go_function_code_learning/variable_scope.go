package main

import "fmt"

var a int = 10

/*
意思就是在func外,每个语句都必须是golang的关键字开始,否则就报这个错，写在函数体内就好。
故此处不可以使用:

	a :=10
*/
func main() {
	fmt.Println("Variable scope code things:")
	a += 10
	fmt.Println(a)
}
