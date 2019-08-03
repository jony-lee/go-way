package main

import "fmt"

type user struct {
	name string
	age  byte
}

type manager struct {
	user
	title string
}

func main() {
	var m manager
	m.name = "tom"
	m.age = 29
	m.title = "CTO"
	fmt.Println(m)
}

//输出：
//{{tom 29} CTO}

//输出顺序是按照结构体排列顺序输出的

//TODO 这里结构体记得是一个结构体内只能嵌入一个，但可以链式嵌入，回头有学到再补充
