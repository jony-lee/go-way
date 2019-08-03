package main

import (
	"fmt"
	"reflect"
)

//类型有一个与之相关的方法集，这决定了它是否实现某个接口
//1、类型T方法集包含所有receiver T方法
//2、类型*T方法集包含所有receiver T + *T方法
//3、匿名嵌入S，T方法集包含所有receiver S方法
//4、匿名嵌入*S，T方法集包含所有receiver S + *S方法
//5、匿名嵌入S或*S，*T方法集包含所有receiver S + *S方法
//理解1
//*T方法集包含receiver T + *T + S + *S方法
//T方法集包含receiver T + S方法
//

//TODO 需要着重理解的知识点
type S struct{}

type T struct {
	S //匿名嵌入字段
	//*S	//如果是*S的话，那么T方法集应包含receiver T + S + *S
}

func (S) sVal()  {}
func (*S) sPtr() {}
func (T) tVal()  {}
func (*T) tPtr() {}

func methodSet(a interface{}) {
	t := reflect.TypeOf(a)

	for i, n := 0, t.NumMethod(); i < n; i++ {
		m := t.Method(i)
		fmt.Println(m.Name, m.Type)
	}
}

func main() {
	var t T
	methodSet(t) //显示T方法集
	fmt.Println("-------------")
	methodSet(&t) //显示*T方法集
}

//TODO go语言学习笔记（雨痕）P135，自己写的代码和书上代码的结果不一致，自己写的没有结果
