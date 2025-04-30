package main

/*
1.Go并发：
	并发是指程序同时执行多个任务的能力。
	Go 语言支持并发，通过 goroutines 和 channels 提供了一种简洁且高效的方式来实现并发。
2.Goroutines：
    Go 中的并发执行单位，类似于轻量级的线程。
    Goroutine 的调度由 Go 运行时管理，用户无需手动分配线程。
    使用 go 关键字启动 Goroutine。
    Goroutine 是非阻塞的，可以高效地运行成千上万个 Goroutine。
3.Channel：
    Go 中用于在 Goroutine 之间通信的机制。
    支持同步和数据共享，避免了显式的锁机制。
    使用 chan 关键字创建，通过 <- 操作符发送和接收数据。
	注释：
		可以使用带缓冲区的通道，称为“通道缓冲区”
		ch:=make(chan int,100)
		如果通道不带缓冲，发送方会阻塞直到接收方从通道中接收了值。
		如果通道带缓冲，发送方则会阻塞直到发送的值被拷贝到缓冲区内；
		如果缓冲区已满，则意味着需要等待直到某个接收方获取到一个值。接收方在有值可以接收之前会一直阻塞。
4.Scheduler（调度器）：
	Go 的调度器基于 GMP 模型，调度器会将 Goroutine 分配到系统线程中执行，并通过 M 和 P 的配合高效管理并发。
    G：Goroutine。
    M：系统线程（Machine）。
    P：逻辑处理器（Processor）。
*/

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	// 定义一个通道，用于传递数据
	fmt.Println("Go goroutines using:")
	// （1）启动goRoutine通道
	go goRoutine()
	// 启动主Routine通道
	for i := 0; i < 10; i++ {
		fmt.Printf("---%v\n", i+1)
		time.Sleep(100 * time.Millisecond)
	}
	// （2）Channel使用
	fmt.Println("Please enter some numbers that use space split it:")
	// 不带缓冲区的通道
	ch := make(chan int)
	noWithBufferChannel(ch)
	// 带缓冲区的通道
	ch1 := make(chan int, 100)
	withBufferChannel(ch1)
}

func goRoutine() {
	// goroutine 是轻量级线程，goroutine 的调度是由 Golang 运行时进行管理的。
	// Go test Routine
	for i := 0; i < 10; i++ {
		fmt.Printf("***%v\n", i+1)
		time.Sleep(100 * time.Millisecond)
	}
}

func noWithBufferChannel(ch chan int) {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	strs := scanner.Text()
	nums := strings.Split(strs, " ")
	// 切片转换为int数组
	arr := make([]int, len(nums))
	for i, v := range nums {
		arr[i] = int(v[0] - '0')
	}
	go sum(arr[:len(arr)/2], ch)
	go sum(arr[len(arr)/2:], ch)
	x, y := <-ch, <-ch
	fmt.Println(x, y, x+y)
}

func withBufferChannel(ch chan int) {
	// 因为ch带缓冲区，且缓冲区长度为100，所以可以先加入几个数字，然后再从通道中取出，
	// 这样就可以避免阻塞，因为缓冲区已满，所以会阻塞直到某个接收方获取到一个值。
	// 接收方在有值可以接收之前会一直阻塞。
	ch <- 1
	ch <- 2
	ch <- 3
	for i := 0; i < 3; i++ {
		fmt.Println(<-ch)
	}
}

func sum(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	c <- sum
}
