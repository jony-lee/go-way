package main

import (
	"fmt"
	"sync"
	"time"
)

//标准库sync提供了互斥和读写锁。
//将Mutex作为匿名字段时，相关方法必须实现pointer-receiver，否则会因为复制导致锁机制失效。

type data struct {
	*sync.Mutex
}

func (d data) test(s string) {
	d.Lock()
	defer d.Unlock()

	for i := 0; i < 5; i++ {
		fmt.Println(s, i)
		time.Sleep(time.Second)
	}
}

func main() {
	var wg sync.WaitGroup
	wg.Add(2)
	var d data
	go func() {
		defer wg.Done()
		d.test("read")
	}()

	go func() {
		defer wg.Done()
		d.test("write")
	}()

	wg.Wait()
}
