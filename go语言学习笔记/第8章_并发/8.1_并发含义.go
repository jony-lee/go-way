package main

import (
	"fmt"
	"time"
)

//并发：逻辑上具备同时处理多个任务的能力
//并行：物理上在同一时刻执行多个并发任务
var c int

func counter() int {
	c++
	return c
}

func main() {
	a := 100

	go func(x, y int) {
		time.Sleep(time.Second) //让goroutine在main逻辑之后执行
		fmt.Println("go:", x, y)
	}(a, counter()) //立即计算并复制参数

	a += 100
	fmt.Println("main:", a, counter())
	time.Sleep(time.Second * 3) //等待goroutine结束，如果不等待，程序结束，goroutine结果无法执行
}
