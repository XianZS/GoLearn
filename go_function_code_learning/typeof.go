package main

import (
	"fmt"
	"reflect"
)

func main() {
	fmt.Println("Go type of something:")
	nums := make([]string, 10)
	fmt.Printf("nums typeof is %v\n", reflect.TypeOf(nums))
	var some []string
	fmt.Printf("some typeof is %v\n", reflect.TypeOf(some))
	number := 1
	fmt.Printf("number typeof is %v\n", reflect.TypeOf(number))
	var x int64
	x = 111
	fmt.Printf("x typeof is %v\n", reflect.TypeOf(x))
}
