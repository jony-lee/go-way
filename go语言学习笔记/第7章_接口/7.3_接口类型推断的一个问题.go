package main

import "fmt"

func main() {
	var x interface{} = 1
	e := x.(int)
	fmt.Println(e)
	//此处类型推断返回两个值，程序没报错
	e1, ok := x.(float64)
	fmt.Println(e1, ok)
	//此处类型推断返回了一个值，程序报错了
	//e2:=x.(float64)
	//fmt.Println(e2)
	//TODO go不是没有重载特性吗，为什么同一个函数既可以返回一个值，又可以返回两个值，底层是怎么实现的？？？
	//利用type switch来进行类型检查
	// type switch 不支持fallthrought
	switch v := x.(type) {
	case nil:
		fmt.Println("nil")
	case *int:
		fmt.Println(*v)
	case func(int) string:
		fmt.Println(v(100))
	case fmt.Stringer:
		fmt.Println(v)
	default:
		fmt.Println("unknown")
	}

}
