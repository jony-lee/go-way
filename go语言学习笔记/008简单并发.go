package main

import (
	"fmt"
	"time"
)

func task(id int) {
	for i := 0; i < 5; i++ {
		fmt.Printf("%d:%d\n", id, i)
		time.Sleep(time.Second * 1)
	}
}

func main() {
	go task(1)
	go task(2)
	time.Sleep(time.Second * 10)
}
