package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Println("Go file code using:")
	var fileName string
	fmt.Println("请输入文件名:")
	_, _ = fmt.Scan(&fileName)
	// [1]创建文件
	//createFile(fileName)
	// [2]打开和关闭文件
	openAndCloseFile(fileName)
	// [3]读取文件内容 = 每次读取一行
	readFileLineMethod(fileName)
	// [4]读取文件内容 = 一次性读取整个文件
	readFileAllMethod(fileName)
	// [5]逐行写入文件
	writeLineFile(fileName)
}

/*
文件操作方式1 >>> 创建文件

	使用os.Create 函数用于创建一个文件，并返回一个 *os.File 类型的文件对象。
	创建文件后，我们通常需要调用 Close 方法来关闭文件，以释放系统资源。
*/
func createFile(fileName string) {
	// 创建文件，如果文件已经存在则会被截断(清空)
	fileObj, _ := os.Create("./" + fileName + ".txt")
	// 确保文件关闭
	defer func(fileObj *os.File) {
		_ = fileObj.Close()
	}(fileObj)
}

/*
文件操作方式2 >>> 文件的打开与关闭

	os.Open 函数用于打开一个文件，并返回一个 *os.File 类型的文件对象。
	打开文件后，我们通常需要调用 Close 方法来关闭文件，以释放系统资源。
*/
func openAndCloseFile(fileName string) {
	// 打开文件
	fileObj, err := os.Open("./" + fileName + ".txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Open file is success!")
	// 打开文件
	defer func(fileObj *os.File) {
		_ = fileObj.Close()
		fmt.Println("文件成功关闭")
	}(fileObj)
}

/*
文件操作方式3 >>> 文件内容的读取

	Go 语言提供了多种读取文件的方式，包括逐行读取、一次性读取整个文件等。
	我们可以使用 bufio 包来逐行读取文件，或者使用 ioutil 包来一次性读取整个文件。
*/
func readFileLineMethod(fileName string) {
	// 打开文件
	fileObj, _ := os.Open("./" + fileName + ".txt")
	defer func(fileObj *os.File) {
		_ = fileObj.Close()
	}(fileObj)
	scanner := bufio.NewScanner(fileObj)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		fmt.Println(err)
	}
}

/*
文件操作方4 >>> 文件内容的读取

	Go 语言提供了多种读取文件的方式，包括逐行读取、一次性读取整个文件等。
	在这个例子中，我们使用 os.ReadFile 一次性读取整个文件的内容，并将其转换为字符串打印出来。
*/
func readFileAllMethod(fileName string) {
	content, _ := os.ReadFile("./" + fileName + ".txt")
	fmt.Println(string(content))
}

/*
文件操作方式5 >>> 逐行写入文件

	Go 语言提供了多种写入文件的方式，包括逐行写入、一次性写入整个文件等。
	在这个例子中，我们使用 bufio.NewWriter 逐行写入文件。
	首先，我们打开文件并创建一个 bufio.Writer 对象。然后，我们逐行写入文件，最后使用 Flush 方法将缓冲区中的内容写入文件。
*/
func writeLineFile(fileName string) {
	fileObj, _ := os.Open("./" + fileName + ".txt")
	defer func(fileObj *os.File) {
		_ = fileObj.Close()
	}(fileObj)
	writer := bufio.NewWriter(fileObj)
	var c int
	fmt.Println("Please input how many line you want to write:")
	_, _ = fmt.Scanln(&c)
	fmt.Println(c)
	scanner := bufio.NewScanner(os.Stdin)
	for i := 0; i < c; i++ {
		scanner.Scan()
		_, err := fmt.Fprintln(writer, scanner.Text())
		if err != nil {
			fmt.Println(err)
			return
		} else {
			/* pass */
		}
		_ = writer.Flush()
	}
}
