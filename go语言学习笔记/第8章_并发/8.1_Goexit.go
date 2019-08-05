package main

import (
	"fmt"
	"runtime"
)

//立刻终止当前任务，运行时确保所有已注册延迟调用被执行。该函数不会影响其他并发任务，不会引发panic。

func main() {
	exit := make(chan struct{})

	go func() {
		defer close(exit)
		defer fmt.Println("a") //执行

		func() {
			defer func() {
				fmt.Println("b", recover() == nil) //执行，recover 返回 nil
			}()
			func() {
				fmt.Println("c")
				runtime.Goexit()       //立即终止整个调用堆栈(终止并依次执行defer，然后退出当前并发)
				fmt.Println("c done.") //不会执行
			}()

			fmt.Println("b done.") //不会执行
		}()

		fmt.Println("a done.") //不会执行

	}()
	<-exit //读nil的channel，阻塞，关闭或写入元素能够解除阻塞
	fmt.Println("main exit.")
}

//如果在main.main里调用GoExit，它会等待其他任务结束，然后让进程崩溃。
