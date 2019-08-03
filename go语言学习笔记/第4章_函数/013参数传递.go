package 第四章_函数

import "fmt"

//不管是指针、引用还是其他类型参数，函数都是值拷贝传递
func test1(x *int) {
	fmt.Printf("pointer:%p,target:%v\n", &x, x)
}

func main() {
	a := 0x100
	p := &a
	fmt.Printf("pointer:%p,target:%v\n", &p, p)
	test1(p)
}

//结果：
//pointer:0xc00008e018,target:0xc000062080
//pointer:0xc00008e028,target:0xc000062080
//可以看到前面的地址是不一样的，但是后面都指向了同一个地址
//即函数创建了一个变量，并把之前的变量中的值（指针地址）赋给了函数的内部变量

//tip如果需要修改传入对象状态的，使用指针类型
//如果不需要修改传入对象状态的，使用值类型
