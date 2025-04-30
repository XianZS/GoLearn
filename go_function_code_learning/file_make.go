package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
)

func main() {
	// 创建文件对象
	fileObj, err := os.Open("./myFile.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	// 读取文件内容
	readerFileBufio(fileObj)
	// 追加文件内容
	writerFileBufio()
	// 关闭文件
	defer func(fileObj *os.File) {
		err := fileObj.Close()
		if err != nil {
			fmt.Println(err)
			return
		} else {
			fmt.Println("文件关闭成功")
		}
	}(fileObj)
}

func readerFileBufio(fileObj *os.File) {
	fmt.Println(fileObj.Name())
	/* 打开文件读取 */
	var index int = 0
	// 创建缓冲区Reader对象
	reader := bufio.NewReader(fileObj)
	for {
		line, err := reader.ReadString('\n')
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Println(err)
			break
		}
		fmt.Print(index+1, ":", line)
		index++
	}
}

func writerFileBufio() {
	fileObj, err := os.OpenFile("./myFileOne.txt", os.O_APPEND|os.O_CREATE, 0666)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer fileObj.Close()
	writer := bufio.NewWriter(fileObj)
	scanner := bufio.NewScanner(os.Stdin)
	var index int = 0
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	for {
		scanner.Scan()
		line := scanner.Text()
		writer.WriteString(line + "\n")
		index += 1
		fmt.Printf("That index is %v\n", index)
		if index >= n {
			break
		}
	}
	writer.Flush()
}
