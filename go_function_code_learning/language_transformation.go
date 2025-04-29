package main

import (
	"fmt"
	"reflect"
	"strconv"
)

/*
	普通强制类型转换
	import "strconv"
	- Strconv make string -> int
		number, err:=strconv.Atoi(string)
	- Strconv make string -> float
		number, err:=strconv.FormatFloat(float64(string), 'f', 小数位数, 64)
	接口类型转换
		- 类型断言：
			接口 -> 其他类型
			类型断言用于将接口类型转换为指定类型，其语法为：value.(type)
			eg: value.(int) 将其转换为int类型
				value.(type) 目标转换类型自行判断
		- 类型转换：
			其他类型 -> 接口
			 类型转换用于将一个接口类型的值转换为另一个接口类型，其语法为：T(value)
*/

func main() {
	fmt.Println("Go a type to b type using:")
	fmt.Println("Please into a string(make of number):")
	var strs string
	_, err := fmt.Scan(&strs)
	if err != nil {
		return
	}
	number, _ := strconv.Atoi(strs)
	// 将string转换为int类型
	fmt.Println("转换之后的int为:", number)
	// 将string转换为float类型
	numberF := strconv.FormatFloat(float64(number), 'f', 3, 64)
	fmt.Println(numberF)
	// 将接口interface类型转换为其他类型(interface to string)
	interfaceToOther()
	// 将其他类型转换为接口interface类型
	otherToInterface()
	// 转换类型自动判断
	printValue("hello")
	printValue(123)
	printValue(false)
}

func interfaceToOther() {
	var i interface{} = "Hello World"
	fmt.Println(i, "type is ", reflect.TypeOf(i))
	str, ok := i.(string)
	if ok {
		fmt.Println(str)
		fmt.Println(str, "type is ", reflect.TypeOf(str))
	}
}

func otherToInterface() {
	// pass
}

func printValue(v interface{}) {
	switch v := v.(type) {
	case int:
		fmt.Println("Integer:", v)
	case string:
		fmt.Println("String:", v)
	case bool:
		fmt.Println("Boolean:", v)
	default:
		fmt.Println("Unknown type")
	}
}
