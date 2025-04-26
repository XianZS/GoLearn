package main

import (
	"bufio"
	"fmt"
	"os"
	"reflect"
	"strings"
)

/*
	切分字符串使用:
		import "strings"
		nums:=strings.Split(被切分字符串,分割字符)
	补充：
		获取变量类型:
			import "reflect"
			some:=reflect.TypeOf(变量) // re:返回变量类型
*/

func main() {
	fmt.Println("split strings")
	scanner := bufio.NewScanner(os.Stdin)
	judgement := scanner.Scan()
	myStr := scanner.Text()
	if !judgement {
		return
	} else {
		fmt.Printf("Now strs is %v\n", myStr)
		strs := strings.Split(myStr, " ")
		fmt.Println(strs, len(strs), reflect.TypeOf(strs[0]), reflect.TypeOf(strs))
	}
}
