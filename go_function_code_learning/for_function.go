package main

import "fmt"

/*
	for i,child := range nums{
		// i meaning index
		// child meaning nums[i]
		fmt.Println(i,child)
	}
*/

func main() {
	fmt.Println("Go for function using things:")
	nums := make([]int, 10)
	for i, child := range nums {
		fmt.Println(i, child)
	}
	fmt.Printf("len(nums)=%v\n", len(nums))
	fmt.Printf("nums is %v\n", nums)
}
