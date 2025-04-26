package main

import "fmt"

/*
	接收number,并且将其-10,输出
*/

func main() {
	fmt.Println("logic judgement function code")
	var number int
	number = 0
	fmt.Printf("Now number is %v\n", number)
	fmt.Print("Please enter a number that type is int:")
	n, err := fmt.Scan(&number)
	if err != nil {
		fmt.Println(err, n)
	} else {
		fmt.Println(err)
		number -= 10
		fmt.Printf("Now number is %v\n", number)
	}
}
