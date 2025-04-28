package main

import (
	"fmt"
)

/*
	type 结构体名 struct {
		变量-1 类型
		变量-2 类型
	}

	实例化接口(结构体):
		book1=bool()
	访问:
		book1.title
*/

type book struct {
	title  string
	author string
	year   int
	bookId int
}

func main() {
	fmt.Println("Go struct codes:")
	// 实例化接口(结构体)
	book1 := book{title: "《天空之城》"}
	fmt.Println(book1)
	// 访问结构体属性
	fmt.Printf("book1.title:%v\n", book1.title)
	// 将结构体作为参数传递
	makeBook(book1)
	fmt.Printf("book1.booId:%v\n", book1.bookId)
}

func makeBook(book1 book) {
	book1.bookId = 1
}
