package main

import "fmt"

//字段名必须唯一、可用“_”补位，支持使用自身指针类型成员。字段名、排列顺序属类型组成部分。除对齐处理外，编译器不会优化调整内部布局
type node struct {
	_    int
	id   int
	next *node
}

type user struct {
	name string
	age  byte
}

func structInit() {

	u1 := user{"Tom", 12} //按顺序初始化

	u2 := user{ //使用字段初始化
		name: "Tom",
		age:  12,
	}

	//u3:=user{"Tom"}				//错误： too few values in user literal
	u4 := user{ //使用命名 初始化,如果有省略字段的化初始化为零值
		name: "Tom",
	}
	fmt.Println(u1, u2, u4)
}

//可以通过指针直接操作结构字段，但不能是多级指针
func structPoint() {
	p := &user{
		name: "Tom",
		age:  20,
	}
	p.name = "Mary"
	p.age++
	//p2 := &p
	//*p2.name = "Jack"	//p2.name undefined (type **user has no field or method name)
}
func main() {
	structInit()
}
