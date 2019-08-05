package 其他

import (
	"fmt"
	"time"
)

func dt() {
	a := time.Now()
	fmt.Println(&a)
	defer func(t time.Time) {
		fmt.Println(time.Now().Sub(t))
	}(time.Now())
	time.Sleep(time.Duration(2) * time.Second)
}
func main() {
	dt()
}

//执行结果
//通过defer延迟执行简单实现获取一段函数的执行时间
//2019-07-29 16:12:52.2150342 +0800 CST m=+0.008008601
//2.0040981s
