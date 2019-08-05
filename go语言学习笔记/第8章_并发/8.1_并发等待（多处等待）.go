package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		wg.Wait()
		fmt.Println("wait exit.")
	}()

	go func() {
		time.Sleep(time.Second)
		fmt.Println("done.")
		wg.Done()
	}()

	wg.Wait()
	fmt.Println("main exit.")
}
