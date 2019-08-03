package main

import "fmt"

//方法和函数一样，除直接调用外，还可赋值给变量，或作为参数传递。
//依照具体引用方式的不同，可分为expression和value两种状态。

//通过类型引用的method expression会被还原为普通函数样式，receiver是第一个参数
//调用时需显示传参。至于类型，可以是T或*T，只要目标方法存在该类型方法集中即可。

type N1 int

func (n N1) test() {
	fmt.Printf("test.n: %p, %d\n", &n, n)
}

func methodExpression() {
	var n N1 = 25
	fmt.Printf("main.n: %p, %d\n", &n, n)
	f1 := N1.test
	f1(n)
	f2 := (*N1).test
	f2(&n)
	fmt.Println("以表达式调用结果")
	//直接以表达式方式调用
	N1.test(n) //receiver作为第一个参数传入
	(*N1).test(&n)
}

//编译器会为method value生成一个包装函数，实现间接调用。至于receiver复制，和闭包的实现方法基本相同。
func methodValue() {
	var n N1 = 100
	p := &n
	n++
	f1 := n.test //此时f1从n=101复制了这个值，并绑定到f1中，因此在f1函数执行时隐式传入了101这个receiver参数

	n++
	f2 := p.test

	n++

	fmt.Printf("main.n: %p, %v\n", p, n)
	f1()
	f2()
}

func main() {
	methodExpression()
	fmt.Println("----------------")
	methodValue()
}
