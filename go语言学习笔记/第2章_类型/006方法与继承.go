package main

import "fmt"

//类似于父类
type user struct {
	name string
	age  byte
}

//类似于子类
type manager struct {
	user
	title string
}

//类似于父类的方法
func (u user) ToString() string {
	return fmt.Sprintf("%+v", u)
}

func main() {
	var m manager
	m.name = "Tom"
	m.age = 29
	m.title = "CTO"
	println(m.ToString())
}
