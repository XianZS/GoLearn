package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println("operator function")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	judgement := strings.Contains(scanner.Text(), "@")
	fmt.Println("judgement:", judgement)
	nums := strings.Split(scanner.Text(), " ")
	fmt.Printf("nums: %v\n", nums)
}
