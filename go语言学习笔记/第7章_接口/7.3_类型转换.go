package main

import "fmt"

type data2 int

func (d data2) String() string {
	return fmt.Sprintf("data:%d", d)
}

//类型推断可将接口变量还原为原始类型，或用来判断是否实现了某个更具体的接口类型
func main() {
	var d data2 = 15
	var x interface{} = d
	if n, ok := x.(fmt.Stringer); ok { //转换为更具体的接口类型
		fmt.Println(n)
	}
	if d2, ok := x.(data2); ok { //转换为原始类型
		fmt.Println(d2)
	}
	e := x.(error) //报错
	fmt.Println(e)
	e, ok := x.(error)
	fmt.Println(e, ok)
	//e,ok := x.(error)
	//fmt.Println(e,ok)
	//TODO 为什么即能接收单值，又能接收两个值呢？
}
