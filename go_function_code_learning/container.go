package main

import "fmt"

/*
	go language's list
	choose-1:
		var nums[size] type
	choose-2:
		var nums=[]type{}
	choose-3:
		nums:=[]type{}
	function using:
		- append(slice切片，添加元素)
			return: 切片对象
*/

func main() {
	fmt.Println("container code things:")
	var nums [10]int
	//for i := 0; i < len(nums); i++ {
	//	nums[i] = i + 1
	//}
	nums[0] = 1
	fmt.Println("len(nums) is", len(nums))
	slice := append(nums[:], 11)
	fmt.Println("len(nums) is", len(nums))
	fmt.Printf("slice is %v,len(slice)=%v\n", slice, len(slice))
	fmt.Println(nums)
}
