package main

import (
	"fmt"
)

func main() {
	fmt.Println("The nil is null")
	var ptr *int
	judgement := ptr == nil
	if judgement {
		fmt.Println("The nil is null")
	}
}
