package main

import "fmt"

func mapAccess() {
	type user struct {
		name string
		age  byte
	}
	m := map[int]user{
		1: {"Tom", 19},
	}
	//m[1].age += 1		//错误：cannot assign to m[1].age
	//1、先把整个元素拎出来，修改再赋值回去
	u := m[1]
	u.age += 1
	m[1] = u
	//2、通过指针的方式原址操作
	m2 := map[int]*user{
		1: &user{"Jack", 20},
	}
	m2[1].age += 1
}

//定义未初始化的字典不能进行写操作，否则会报错，但可以进行读操作。
func mapWrite() {
	var m map[string]int
	fmt.Println(m["a"])
	//m["a"] = 1		//panic: assignment to entry in nil map
}

//迭代期间删除或新增键值是安全的
//但不能保证迭代操作删除新增的键值
func mapOpt() {
	m := make(map[int]int)
	for i := 0; i < 10; i++ {
		m[i] = i + 10
	}
	for k := range m {
		if k == 5 {
			m[100] = 1000
		}
		delete(m, k)
		fmt.Println(k, m)
	}
}
func main() {
	mapOpt()
}

//TODO 字典理解比较困难，尤其是这里安全的，还有一些并发是什么情况需要再学习
