package main

import (
	"fmt"
)

/*
对于slice类型：
	1.len()
		re当前元素个数
	2.cap()
		re当前slice的最大长度
	3.[]切片处理
		re[beginIndex,endIndex)
	4.append()函数
		支持多元素同时添加
		append(slice,child1,child2,child3,……,childN)
	5.copy()函数
		copy(目标,被使用)
*/

func main() {
	fmt.Println("Go slice using:")
	// 初始化
	nums := make([]int, 3, 10)
	for index, _ := range nums {
		nums[index] = index + 1
	}
	fmt.Printf("nums:%v\n", nums)
	// 求len
	lenNums, judgement1 := getLen(nums)
	// 求cap
	capNums, judgement2 := getCap(nums)
	if judgement1 {
		fmt.Printf("len:%v\n", lenNums)
	} else {
		fmt.Println("Error")
	}
	if judgement2 {
		fmt.Printf("cap:%v\n", capNums)
	} else {
		fmt.Println("Error")
	}
	// 求切片
	reNums := getChild(nums)
	fmt.Printf("reNums:%v\n", reNums)
	// append添加元素
	appendChild(nums)
	// copy()复制nums's slice
	copyNums()
}

// 求当前元素个数
func getLen(nums []int) (int, bool) {
	if nums == nil {
		return 0, false
	} else {
		return len(nums), true
	}
}

// 求slice的最大长度
func getCap(nums []int) (int, bool) {
	if nums == nil {
		return 0, false
	} else {
		return cap(nums), true
	}
}

// []切片处理
func getChild(nums []int) []int {
	return nums[2:8]
}

// append()添加元素处理
func appendChild(nums []int) {
	newNums := append(nums, 1, 2, 3, 4)
	fmt.Println(newNums)
}

// copy()复制元素处理
func copyNums() {
	nums := make([][10]int, 10)
	for x, _ := range nums {
		for y, _ := range nums[x] {
			nums[x][y] = x + y
		}
	}
	newNums := make([][10]int, 10)
	copy(newNums, nums)
	newNums[0][0] = -100
	fmt.Println(newNums[0][0], nums[0][0])
}
