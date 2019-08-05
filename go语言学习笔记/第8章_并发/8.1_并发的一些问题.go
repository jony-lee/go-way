package main

import (
	"fmt"
	"sync"
)

//与线程不同，goroutine任务无法设置优先级，无法获取编号
//没有局部存储（TLS），甚至连返回值都会被抛弃（如需要处理返回值，使用channel）
//TODO 不理解，“如使用map作为局部存储容器，建议做同步处理，因为允许时会对其做并发读写检查”
func main() {
	var wg sync.WaitGroup
	var gs [5]struct {
		id     int //实现goroutine的编号
		result int //实现goroutine的返回值或局部存储
	}

	for i := 0; i < len(gs); i++ {
		wg.Add(1)

		go func(id int) {
			defer wg.Done()
			gs[id].id = id
			gs[id].result = (id + 1) * 100
		}(i)

	}

	wg.Wait()
	fmt.Printf("&+v\n", gs)

}
