package main

import "fmt"

type data1 struct{}

func (data1) string() string {
	return "data1"
}

type node struct {
	data interface { //匿名接口类型
		string() string
	}
}

func main() {
	var t interface { //定义匿名接口变量
		string() string
	} = data1{} //实例化接口

	n := node{
		data: t,
	}
	fmt.Println(n.data.string())
}
