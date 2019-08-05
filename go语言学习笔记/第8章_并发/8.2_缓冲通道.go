package main

import (
	"fmt"
	"unsafe"
)

func channelBuf() {
	c := make(chan int, 3)
	c <- 1
	c <- 2
	fmt.Println(<-c)
	fmt.Println(<-c)
}

func channelEqual() {
	var a, b = make(chan int, 3), make(chan int)
	var c chan bool
	//判断通道是否为同一对象或nil
	fmt.Println(a == b)
	fmt.Println(c == nil)
	fmt.Printf("%p,%d\n", a, unsafe.Sizeof(a))

	//判断通道缓冲区大小和当前已缓冲数量
	a <- 1
	fmt.Println("异步通道 a:", len(a), cap(a))
	fmt.Println("同步通道 b:", len(b), cap(b))

}

func main() {
	//channelBuf()
	channelEqual()
}
