package main

import (
	"fmt"
)

func main() {
	fmt.Println("if-else line using:")
	var judgement bool
	scan, err := fmt.Scan(&judgement)
	if err != nil {
		return
	} else {
		if scan == 1 {
			if judgement {
				fmt.Println("Successful")
			} else {
				fmt.Println("Failed")
			}
		} else {
			fmt.Println("Error")
		}
	}
}
