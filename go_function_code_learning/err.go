package main

import (
	"errors"
	"fmt"
)

func main() {
	fmt.Println("Go err code using:")
	err := errors.New("This is an error")
	fmt.Println(err)
}
