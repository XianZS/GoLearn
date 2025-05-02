package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	fmt.Println("Go file code using:")
	fmt.Println("请输入文件名:")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	fileName := scanner.Text()
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
	// [6]删除文件
	deleteFile()
	// [7]获取文件信息
	getFileInformation(fileName)
	// [8]创建目录
	createDir()
	// [9]读取目录下的内容
	readDir()
	// [10]删除目录
	deleteDir()
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
	fileObj, _ := os.OpenFile("./"+fileName+".txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)
	defer func(fileObj *os.File) {
		_ = fileObj.Close()
	}(fileObj)
	writer := bufio.NewWriter(fileObj)
	fmt.Printf("Please input how many line you want to write:")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	c, _ := strconv.Atoi(scanner.Text())
	for i := 0; i < c; i++ {
		fmt.Println("please input line:")
		scanner.Scan()
		lineObj := scanner.Text()
		fmt.Println(lineObj)
		fprintln, err := fmt.Fprintln(writer, lineObj)
		if err != nil {
			fmt.Println(fprintln)
			return
		}
		err2 := writer.Flush()
		if err2 != nil {
			fmt.Println(err2)
			return
		}
	}
}

/*
文件操作方式6 >>> 删除文件

	Go 语言提供了多种删除文件的方式，包括删除单个文件、删除目录等。
	在这个例子中，我们使用 os.Remove 函数删除文件。
	首先，我们创建一个文件，然后使用 os.Remove 函数删除文件。
	如果删除文件时出错，我们打印错误信息并返回。否则，我们打印删除文件成功的信息。
*/
func deleteFile() {
	fileObj, _ := os.Create("./test.txt")
	_ = fileObj.Close()
	err := os.Remove("./test.txt")
	if err != nil {
		fmt.Println(err)
		return
	} else {
		fmt.Println("Delete text.txt success!")
	}
}

/*
文件操作方式7 >>> 文件信息操作
文件名:  myFileName.txt
文件大小:        100
文件权限:        -rw-rw-rw-
文件修改时间:    2025-05-02 13:11:35.2489903 +0800 CST
文件是否是目录:  false
文件系统:        &{32 {343202337 31177168} {2901580719 31177504} {2901580719 31177504} 0 100}
*/
func getFileInformation(fileName string) {
	fileObj, _ := os.Stat("./" + fileName + ".txt")
	fmt.Println("文件名:\t", fileObj.Name())
	fmt.Println("文件大小:\t", fileObj.Size())
	fmt.Println("文件权限:\t", fileObj.Mode())
	fmt.Println("文件修改时间:\t", fileObj.ModTime())
	fmt.Println("文件是否是目录:\t", fileObj.IsDir())
	fmt.Println("文件系统:\t", fileObj.Sys())
}

/*
文件操作方式8 >>> 创建目录

	Go 语言提供了多种创建目录的方式，包括创建单个目录、创建多级目录等。
	在这个例子中，我们使用 os.Mkdir 函数创建目录。
	首先，我们创建一个目录，然后使用 os.Mkdir 函数创建目录。
	如果创建目录时出错，我们打印错误信息并返回。否则，我们打印创建目录成功的信息。
*/
func createDir() {
	// 创建单个目录
	err := os.Mkdir("testDir", 0755)
	if err != nil {
		fmt.Println(err)
		return
	}
	// 创建多级目录
	err = os.MkdirAll("testDir/testDir2", 0755)
	if err != nil {
		fmt.Println(err)
		return
	}
}

/*
文件操作方式9 >>> 读取目录下的内容
conference.go
container.go
err.go
file_make.go
file_using.go
for_function.go
func.go
goroutines.go
hello.go
ifelse.go
iii.txt
input.go
interface.go
lambda_function.go
language_transformation.go
logic_judgement.go
map.go
myFile.txt
myFileName.txt
myFileOne.txt
operator.go
ptr_nil.go
recursion.go
revalues.go
slice.go
split_strings.go
struct.go
switch.go
testDir
type_switch.go
typeof.go
variable.go
variable_scope.go
*/
func readDir() {
	// 读取目录
	dirObj, _ := os.ReadDir("./")
	for _, entry := range dirObj {
		info, _ := entry.Info()
		fmt.Println(info.Name())
	}
}

/*
文件操作方式10 >>> 删除目录
*/
func deleteDir() {
	err := os.RemoveAll("./testDir")
	if err != nil {
		fmt.Println(err)
		return
	}
	err = os.RemoveAll("./testDir/testDir2")
	if err != nil {
		fmt.Println(err)
		return
	}
}
