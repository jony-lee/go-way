package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		//建议在goroutine外执行Add累加操作，避免Add尚未执行，Wait已经退出。
		wg.Add(1) //计数增加
		go func(id int) {
			defer wg.Done() //计数减少
			time.Sleep(time.Second * 1)
			fmt.Println("goroutine", id, "done.")
		}(i)
	}
	fmt.Println("main...")
	wg.Wait() //等待计数清空
	fmt.Println("main exit.")
}
