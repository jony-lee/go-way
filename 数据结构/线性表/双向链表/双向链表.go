package main

import "fmt"

//TODO
type List struct {
	data interface{}
	a    int
	pre  *List
}

func main() {
	var a *List
	fmt.Println(a.data)
}
