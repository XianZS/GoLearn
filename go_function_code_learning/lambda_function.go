package main

import (
	"fmt"
)

func main(){
	fmt.Println("Go lambda using:")
	add:=func(a int,b int)int{
		return a+b
	}
	number:=add(1,2)
	fmt.Println(number)
}