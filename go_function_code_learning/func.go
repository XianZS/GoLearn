package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println("Go function using:")
	nums := getNums("jom")
	fmt.Printf("nums: %v\n", nums)
}

func getNums(name string) []string {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	fmt.Printf("name:%v\n", name)
	nums := strings.Split(scanner.Text(), " ")
	fmt.Printf("nums:%v\n", nums)
	return nums
}

func setNums(nums []string) {
	// pass
}
