package main

import "fmt"

func main() {
	data := [3]int{10, 20, 30}
	for i, x := range data { //从data的复制品中取值
		if i == 0 {
			data[0] += 100
			data[1] += 200
			data[2] += 300
		}
		fmt.Printf("x:%d,data:%d\n", x, data[i])
	}
	data = [3]int{10, 20, 30}
	for i, x := range data[:] { //仅复制slice，不包括底层array
		if i == 0 {
			data[0] += 100
			data[1] += 200
			data[2] += 300
		}
		fmt.Printf("x:%d,data:%d\n", x, data[i])
	}
}
