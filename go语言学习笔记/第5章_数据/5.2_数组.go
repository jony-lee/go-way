package main

import "fmt"

//数组声明方式
func arrayDefine() {

	var a [1]int
	b := [2]int{1, 1}
	c := [3]int{1, 1: 2}      //指定索引位置3初始化值为10
	d := [...]int{1, 2, 3, 4} //按初始化值数量确定数组长度
	e := [...]int{1, 4: 4}    //支持索引初始化，但要注意其长度
	fmt.Println("一般类型初始化")
	fmt.Println(a, b, c, d, e)
	fmt.Println(len(a), len(b), len(c), len(d), len(e))

	//对于结构体等复合类型，可省略元素初始化类型标签
	type user struct {
		name string
		age  byte
	}
	f := [...]user{
		{"tom", 20},
		{"mary", 18},
	}
	fmt.Println("结构体类型初始化")
	fmt.Println(f)
}

//多维数组定义
func arrayDefineMulti() {
	a := [2][2]int{
		{1, 2},
		{3, 4},
	}
	b := [...][2]int{
		{10, 20},
		{30, 40},
	}
	c := [...][2][2]int{ //只能在第一维使用[...]
		{
			{1, 2},
			{3, 4},
		},
		{
			{10, 20},
			{30, 40},
		},
	}
	fmt.Println()
	fmt.Println("多维数组定义")
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
}

func main() {
	arrayDefine()
	arrayDefineMulti()
}
