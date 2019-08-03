package main

func main() {
	x, y := 1, 2
	x, y = y+4, x+10 //多变量赋值时，先全部计算右边式子再赋值
	println(x, y)
}
