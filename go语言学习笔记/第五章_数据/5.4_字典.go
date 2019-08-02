package main

import "fmt"

//初始化
//字典内存储顺序为乱序存储
func mapInit(){
	m:= make(map[string]int)
	m["a"] = 1
	m["b"] = 2

	m1:= map[string]int{"a":1,"b":2}

	m2:=map[int]struct{
		x int
	}{
		1:{x:100},
		2:{x:200},
	}
	fmt.Println(m,m1,m2)

	//map的删除操作
	delete(m,"a")
	fmt.Println(m)

	//访问不存在的键值时，默认返回零值，不会引发错误。
	a:=m["c"]
	fmt.Println(a)
	//使用ok-idiom模式判断key是否存在
	if v,ok :=m["c"];ok{
		fmt.Println(v)
	}
}

func main() {
	mapInit()
}