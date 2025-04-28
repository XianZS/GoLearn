package main

import "fmt"

/*
	返回多个数值
	func function_name() (default_type,default_type){
		// code......
	}
*/

func main() {
	fmt.Println("return values example x,y")
	x, y := re()
	fmt.Printf("x:%v\n", x)
	fmt.Printf("y:%v\n", y)
}

func re() (string, string) {
	return "x", "y"
}
