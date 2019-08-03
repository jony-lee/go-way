package main

import "fmt"

//匿名字段，没有名字，仅有类型的字段，也被称为嵌入字段或嵌入类型
type attr struct {
	perm int
}
type file struct {
	name string
	attr //仅有类型名
}

func structAnonymousField() {
	f := file{
		name: "test.dat",
		attr: attr{ //从编译器角度看，只是隐式第以类型名字作为字段名字。
			perm: 0775, //可直接引用匿名字段成员，但初始化的时候必须当作独立字段
		},
	}
	f.perm = 0644       //直接设置匿名字段成员
	fmt.Println(f.perm) //直接读取匿名字段成员
}

func main() {
	structAnonymousField()
}
