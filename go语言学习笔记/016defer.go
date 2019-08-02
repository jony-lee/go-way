package main

import (
	"sync"
	"testing"
)

//误用，不要在main函数中使用defer，加入在main函数中开启一个文件，然后使用defer进行操作，这是，defer会在main函数结束时才关闭资源，于是期间文件资源就一直开着
//慎用，延迟调用会损耗性能
var m sync.Mutex
func call(){
	m.Lock()
	m.Unlock()
}
func deferCall(){
	m.Lock()
	defer m.Unlock()
}

func BenchamarkCall(b *testing.B){
	for i := 0;i<b.N;i++{
		call()
	}
}
func BenchamarkDefer(b *testing.B){
	for i := 0;i<b.N;i++{
		deferCall()
	}
}