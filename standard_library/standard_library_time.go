package main

import (
	"fmt"
	"time"
)

func main() {
	nowTime := time.Now()
	// [1]输出当前时间
	fmt.Println(nowTime)
	fmt.Println(nowTime.Year(), nowTime.Month(), nowTime.Day())
	fmt.Println(nowTime.Hour(), nowTime.Minute(), nowTime.Second())
	// [2]输出当前时间的Unix时间戳
	fmt.Println(nowTime.Unix())
	// [3]格式化当前时间
	fmt.Println(nowTime.Format("2006/01/02"))
}
