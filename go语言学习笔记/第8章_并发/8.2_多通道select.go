package main

import (
	"fmt"
	"sync"
)

//select 随机执行一个可运行的 case，select不会选择发生阻塞的通道。
//当所有通道都不可用时，select会执行default语句。如此可避开select阻塞，但需注意处理外层循环，以免陷入空耗。
func main() {
	var wg sync.WaitGroup
	wg.Add(2)
	a, b := make(chan int), make(chan int)
	go func() {
		defer wg.Done()

		for {
			var (
				name string
				x    int
				ok   bool
			)
			select { //随机选择可用channel接收数据
			case x, ok = <-a:
				name = "a"
			case x, ok = <-b:
				name = "b"

			}
			if !ok { //如果任一通道关闭，则终止接收
				return
			}
			fmt.Println(name, x)
		}
	}()

	go func() {
		defer wg.Done()
		defer close(a)
		defer close(b)

		for i := 0; i < 10; i++ {
			select { //随机选择发送channel
			case a <- i:
			case b <- i * 10:

			}
		}
	}()
	wg.Wait()

}
