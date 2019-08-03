package main

import "fmt"

//空结构，这类长度为零的对象通常都指向runtime.zerobase变量
//空结构可以作为通道元素类型，用于事件通知

func structEmpty() {
	a := struct{}{}
	b := [10]struct{}{}
	c := b[:]
	d := [0]int{}
	fmt.Printf("%p,%p,%p,%p\n", &a, &b[0], &c[0], &d)
}

func structEvent() {
	exit := make(chan struct{})
	go func() {
		fmt.Println("Hello world!")
		exit <- struct{}{}
	}()
	<-exit
	fmt.Println("end.")

}
func main() {
	structEvent()
}
