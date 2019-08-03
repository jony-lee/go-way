package main

import "fmt"

//方法相当于有归属者的函数
//可以为当前包，以及除接口和指针以外的任何类型定义方法
//方法同样不支持重载
//如方法内部并不引用实例，可省略参数名，仅保留类型
type N int

func (n N) toString() string {
	return fmt.Sprintf("%#x", n)
}
func (N) sayHello() { //该处方法内部不引用实例，因此省略参数名n
	fmt.Println("hello")
}

func (n N) value() {
	n++
	fmt.Printf("v: %p, %v\n", &n, n)
}

func (n *N) pointer() {
	(*n)++
	fmt.Printf("p: %p, %v\n", n, *n)
}

//调用方法
func methodCall() {
	var a N = 25
	fmt.Println(a.toString())
	a.sayHello()
	fmt.Printf("a: %p, %v\n", &a, a)
	a.value()   //通过值改变value的时候地址发生了变化
	a.pointer() //通过指针改变value的时候地址没有发生变化
	fmt.Printf("a: %p, %v\n", &a, a)
}

//不能用多级指针调用方法
//TODO 疑惑，多级指针能做哪些事情，不能做哪些事情

//如何选择方法的receiver（上述的方法中的n）的类型
//1、要修改实力状态，用*T
//2、无须修改状态的小对象或固定值，建议用T
//3、大对象建议用*T，以减少复制成本
//4、引用类型、字符串、函数等指针包装对象，直接用T      ？？？？
//5、若包含Mutex等同步字段的用*T，避免因复制造成锁操作无效
//6、其他无法确定的情况，都用*T
func main() {
	methodCall()
}
