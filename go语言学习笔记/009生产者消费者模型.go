package main

import "time"

//消费者
func consumer(data chan int, done chan bool) {
	for x := range data {
		println("receive:", x)
	}
	done <- true //通知main，消费结束
}

//生产者
func producter(data chan int) {
	for i := 0; i < 4; i++ {
		println("product", i)
		data <- i
		time.Sleep(time.Second)
	}
	close(data)
}

func main() {
	done := make(chan bool)
	data := make(chan int)
	go consumer(data, done)
	go producter(data)
	<-done //阻塞，直到消费者发回结束信号
}
