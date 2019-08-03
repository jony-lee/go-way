package main

import "fmt"

//const(
//	x = iota
//	y
//	z
//)
//const(
//	_ = iota
//	KB = 1<<(10*iota)
//	MB
//	GB
//)
//const(
//	_,_ = iota,iota*10
//	a,b
//	c,d
//)
const (
	r = 789
	a = iota
	b
	c = 100
	d
	e = iota
	f
)

//const(
//	a = iota
//	b float32 = iota
//	c = iota
//)
//type color byte
//
//const(
//	black color = iota
//	red
//	blue
//)

//iota其实就是代表const中这一行的行号
//能够自动延续上一行的iota，但中断后不能延续

//不赋值其实就是 默认复制上一行表达式
//iota是变化的，也因此就可以通过iota来生成一定规律常量

func main() {
	fmt.Println(a, b, c, d, e, f)
}
