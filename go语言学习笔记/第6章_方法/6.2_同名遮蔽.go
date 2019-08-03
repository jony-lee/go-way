package main

import "fmt"

type user struct{}

type manager struct {
	user
}

func (user) toString() string {
	return "user"
}
func (m manager) toString() string {
	return "manager"
}

func main() {
	var m manager
	fmt.Println(m.toString())
	fmt.Println(m.user.toString())
}

//manager本身的方法遮蔽了user的方法，有点像是面向对象语言中的重写，即方法覆盖
//这种特性就类似于面向对象语言中的继承和方法重写
