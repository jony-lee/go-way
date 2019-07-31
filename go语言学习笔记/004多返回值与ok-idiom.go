package main

import "fmt"

func main() {
	m := make(map[string]int)
	m["a"] = 1
	x, ok := m["a"]
	fmt.Println("存在元素:", x, ok)
	x, ok = m["b"]
	fmt.Println("不存在元素:", x, ok)
	delete(m, "a") //删除元素

}

//结果：
//存在元素: 1 true
//不存在元素: 0 false
//多返回值中设置一个名为ok的变量来判断操作是否成功
//如果成功ok值为true，如果失败ok值为false
//tips：因为很多操作默认返回零值，因此需要额外说明
//例如：_,err := doSomeThing()，如果成功了err为nil，如果失败了会得到相应的error
