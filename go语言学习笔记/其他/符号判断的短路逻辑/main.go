package main

import "fmt"

func main() {

	//运行结果：a
	if true || b() { //go 的 if 判断采用短路求值，值已经确定后续的表达式不会计算也不会被调用
		fmt.Println("a")
	}

	//运行结果：ba
	if or(true, b()) { //函数的参数必须被即时求值的
		fmt.Println("a")
	}

	//运行结果：a
	if orfunc(true, b) { //传入函数参数的函数实现惰性求值
		fmt.Print("a")
	}
}

func b() bool {
	fmt.Print("b")
	return false
}

func or(a, b bool) bool {
	return a || b
}

func orfunc(a bool, b func() bool) bool {
	return a || b()
}
