package 第四章_函数

import (
	"log"
	"time"
)

//多参数函数建议使用结构体来进行参数传递
type serverOption struct {
	address string
	port    int
	path    string
	timeout time.Duration
	log     *log.Logger
}

func newOption() *serverOption {
	return &serverOption{
		address: "0.0.0.0",
		port:    8080,
		path:    "/var/test",
		timeout: time.Second * 5,
		log:     nil,
	}
}

func server(option *serverOption) {}

func main() {
	opt := newOption()
	opt.port = 8085
	server(opt)
}
