package main

import (
	"bufio"
	"fmt"
	"os"
)

/*
fmt.input:
	- fmt.Scan(&参数1,&参数2)
		可以换行输入，自动省略换行符，以空格或回车作为分割符
	- fmt.Scanf("%v,%v",参数1,参数2)
		必须在一行输入，以空格作为分割符
*/

func main() {
	fmt.Println("input code things")
	var name string
	var age int
	fmt.Scanf("%s%d", &name, &age)
	fmt.Printf("&name:%v\n", &name)
	fmt.Printf("name:%v\n,age:%v\n", name, age)
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Enter a line things: ")
	scanner.Scan()
	line := scanner.Text()
	fmt.Println(line)
	fmt.Println("***")
}
