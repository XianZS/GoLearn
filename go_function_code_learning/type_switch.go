package main

import (
	"fmt"
)

func main() {
	fmt.Println("Go type switch using:")
	printType(123)
	printType("123")
	printType(false)
	printType([]int{})
}
func printType(val interface{}) {
	switch val.(type) {
	case int:
		fmt.Println("Int")
	case string:
		fmt.Println("String")
	case bool:
		fmt.Println("Bool")
	default:
		fmt.Println("UnKnow")
	}
}
