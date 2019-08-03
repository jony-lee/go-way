package main

import "fmt"

type say interface {
	say() string
}

type cater interface {
	say
	run()
}
type tom struct{}

func (*tom) run() {}
func (tom) say() string {
	return ""
}
func pp(a say) {
	fmt.Println(a.say())
}

func main() {
	var d tom
	var t cater = &d
	pp(t) //隐式地转换为子集接口

	var s say = t //超集接口变量可以转化为子集接口变量，反之则不行
	//因为实现超集接口所有方法，则表明实现了其子集的所有方法
	//反之，实现了子集所有方法则不一定实现了所有超集接口方法

	fmt.Println(s.say())

	t.run()
	fmt.Println(t.say())
}
