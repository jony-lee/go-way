package main

import "fmt"

func test(x int) func() { //返回函数类型
	return func() { //匿名函数
		fmt.Println(x) //闭包
	}
}

func main() {
	x := 100
	f := test(x)
	f()
}

//输出
//100
