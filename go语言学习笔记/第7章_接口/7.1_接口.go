package main

import "fmt"

//Go实现接口机制很简洁，只要目标类型方法集内包含接口声明的全部方法，就被认为实现了接口。
//TODO 实现接口有什么用呢？？
//接口
//1、不能有字段
//2、不能定义自己的方法
//3、只能声明方法、不能实现
//4、可嵌入其他接口类型
//总结，接口相当于定义了一套规范

type tester interface {
	test()
	string() string
}
type data struct{}

func (*data) test()         {}
func (data) string() string { return "" }

func interfaceInit() {
	var d data
	var t tester = d //data does not implement tester 		这里涉及到方法集的问题
	// (test method has pointer receiver)
	//var t tester = &d
	t.test()
	fmt.Println(t.string())
}
func main() {
	interfaceInit()
}
