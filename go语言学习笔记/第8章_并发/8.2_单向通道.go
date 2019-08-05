package main

import (
	"fmt"
	"sync"
)

//通道默认时双向的，单我们可以限制收发操作的方向来获取更严谨的操作逻辑
func main() {
	var wg sync.WaitGroup
	wg.Add(2)

	c := make(chan int)
	var send chan<- int = c //使用类型转换来获取单向通道
	var recv <-chan int = c
	fmt.Println(&c, &send, &recv)
	go func() {
		defer wg.Done()
		for x := range recv {
			fmt.Println(x)
		}
	}()

	go func() {
		defer wg.Done()
		defer close(c)

		for i := 0; i < 3; i++ {
			send <- i
		}
	}()
	wg.Wait()
}
