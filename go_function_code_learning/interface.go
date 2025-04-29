package main

import (
	"fmt"
)

/*
1.接口（interface）是 Go 语言中的一种类型，用于定义行为的集合，它通过描述类型必须实现的方法，规定了类型的行为契约。
Go 语言提供了另外一种数据类型即接口，它把所有的具有共性的方法定义在一起，任何其他类型只要实现了这些方法就是实现了这个接口。
Go 的接口设计简单却功能强大，是实现多态和解耦的重要工具。
接口可以让我们将不同的类型绑定到一组公共的方法上，从而实现多态和灵活的设计。
2.接口类型变量：
接口变量可以存储实现该接口的任意值。
接口变量实际上包含了两个部分：
动态类型：存储实际的值类型。
动态值：存储具体的值。
3.特例
零值接口：
	接口的零值是 nil。
	一个未初始化的接口变量其值为 nil，且不包含任何动态类型或值。
空接口：
	定义为 interface{}，可以表示任何类型。
4.接口的常见用法
	多态：不同类型实现同一接口，实现多态行为。
	解耦：通过接口定义依赖关系，降低模块之间的耦合。
	泛化：使用空接口 interface{} 表示任意类型。
5.class的组成
	在go语言中，一个class类对象由interface、struct和function组成。
	其中interface指明了class类对象所具有的方法
	struct指明了class类对象所具有的属性字段
	function具体实现了class类对象的interface定义的方法
6.相关语法格式如下
	interface:
		type 接口名 interface{
			function-1 return_values_type_1
			function-2 return_values_type_2
		}
	struct:
		type struct-name struct{
			属性字段-1 属性字段-type
			属性字段-2 属性字段-type
		}
	function:
		func (结构体对象 结构体类型) interface魔术方法 (return-type){
			// 方法体
		}
*/

// 先定义一个接口People接口
type people interface {
	getName() string
	getAge() int
	getAddress() string
	setName(name string) bool
	setAge(age int) bool
	setAddress(address string) bool
}

// Student 定义一个 Student 结构体
type Student struct {
	name    string
	age     int
	address string
}

func (stu Student) getName() string {
	return stu.name
}

func (stu Student) getAge() int {
	return stu.age
}

func (stu Student) getAddress() string {
	return stu.address
}

func (stu Student) setName(name string) bool {
	stu.name = name
	return true
}

func (stu Student) setAge(age int) bool {
	stu.age = age
	return true
}

func (stu Student) setAddress(address string) bool {
	stu.address = address
	return true
}

func main() {
	fmt.Println("Go interface using:")
	stu := Student{name: "jom", age: 123, address: "ShangHai"}
	fmt.Println(stu)
	fmt.Println(stu.getName())
	fmt.Println(stu.getAge())
	fmt.Println(stu.getAddress())
	stu.setName("kom")
	stu.setAge(123)
	stu.setAddress("BeiJing")
	fmt.Println(stu.getName(), stu.getAge(), stu.getAddress())
}
