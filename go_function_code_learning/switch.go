package main

import (
	"bufio"
	"fmt"
	"os"
)

/*
go语言的switch语句：

	go语言的switch-case自带break，如果需要使用到其他地方case，则需要加fallthrough
	switch cho{
	case cho1:
		// code - 1
	case cho2:
		// code - 2
	default:
		// code - default
	}
*/
func main() {
	fmt.Println("Go switch learning:")
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Printf("Please inter your city: ")
	scanner.Scan()
	city := scanner.Text()
	switch city {
	case "ShangHai":
		fmt.Printf("You choose city is %v", city)
		fallthrough
	case "BeiJing":
		fmt.Printf("You choose city is %v", city)
	case "NanJing", "HangZhou":
		fmt.Printf("You choose city is %v", city)
	default:
		fmt.Println("This is DEFAULT!")
	}
}
