package main

import "fmt"

//Go实现接口机制很简洁，只要目标类型方法集内包含接口声明的全部方法，就被认为实现了接口。
//因为要通过接口类型调用实例化方法，如果有接口方法没实现，就没法调用这个方法，也就不算是实现了接口。
//TODO 实现接口有什么用呢？？
//接口
//1、不能有字段
//2、不能定义自己的方法
//3、只能声明方法、不能实现
//4、可嵌入其他接口类型
//总结，接口相当于定义了一套规范
//接口通常以er作为名称后缀，方法名是声明组成部分，但参数名可不同或省略

type tester interface {
	test()
	string() string
}
type data struct{}

func (*data) test()         {}
func (*data) dataTest()     {}
func (data) string() string { return "" }

func interfaceInit() {
	var d data
	//var t tester = d //data does not implement tester 		这里涉及到方法集的问题
	// (test method has pointer receiver)
	var t tester = &d
	t.test() //接口实现，可通过接口类型变量直接调用实例化接口方法了
	fmt.Println(t.string())
}

//声明一个空接口，用途类似于面向对象里的根类型Object，可被赋值为任何类型的对象。
//type emptier interface {}

func main() {
	interfaceInit()
}
