package main

import (
	"fmt"
	"math"
	"runtime"
	"sync"
)

//线程数默认与处理器数相等，可通过runtime.GOMAXPROCS函数修改

//测试目标函数
func count() {
	x := 0
	for i := 0; i < math.MaxUint32; i++ {
		x += i
	}
	fmt.Println(x)
}

//循环执行
func test(n int) {
	for i := 0; i < n; i++ {
		count()
	}
}

//并发执行
func test2(n int) {
	var wg sync.WaitGroup
	wg.Add(n)
	for i := 0; i < n; i++ {
		go func() {
			count()
			wg.Done()
		}()
	}
	wg.Wait()
}

func main() {
	n := runtime.GOMAXPROCS(0)
	test2(n)
}
