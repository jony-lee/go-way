package main

import "sync"

//可以像访问匿名字段成员那样调用其他方法，由编译器负责查找
type data struct {
	sync.Mutex
	buf [1024]byte
}

func main() {
	d := data{}
	d.Lock()
	defer d.Unlock()
}
