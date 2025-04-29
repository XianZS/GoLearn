package main

import "fmt"

func main() {
	fmt.Println("Go recursion using:")
	var n uint
	_, err := fmt.Scan(&n)
	if err != nil {
		fmt.Println(err)
		return
	}
	// 求阶乘
	fmt.Printf("%v! is %v\n", n, factorial(n))
	// 求斐波拉契数列
	fmt.Printf("Fibonacci's %v is :%v\n", n, fibonacci(n))
}

func factorial(x uint) uint {
	if x == 0 {
		return 1
	} else {
		return x * factorial(x-1)
	}
}

func fibonacci(n uint) uint {
	if n < 2 {
		return n
	}
	return fibonacci(n-1) + fibonacci(n-2)
}
