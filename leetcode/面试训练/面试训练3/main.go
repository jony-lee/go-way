package main

import "fmt"

/*
【题目】
给定由若干 0 和 1 组成的数组 A。我们定义 N_i：从 A[0] 到 A[i] 的第 i 个子数组被解释为一个二进制数（从最高有效位到最低有效位）。

返回布尔值列表 answer，只有当 N_i 可以被 5 整除时，答案 answer[i] 为 true，否则为 false。

【思路】
//TODO



*/

//【代码】

func prefixesDivBy5(A []int) []bool {
	ret := make([]bool, len(A))
	var num int
	for i, v := range A {
		num = num<<1 + v
		if num%5 == 0 {
			ret[i] = true
		}
		num %= 5
	}
	return ret
}

//【主函数】
func main() {
	a := []int{1, 1, 0, 1, 1, 1, 1, 1, 1, 1, 1, 0, 1, 1, 1, 0, 0, 1, 0}
	fmt.Println((prefixesDivBy5(a)))
}

/*
【总结】：






*/
//[false,false,true,false,false,false,false,false,false,false,true,true,true,true,true,true,false,false,false,false,false,false,false,false,false,false,false,false,false,false,false,false,false,false,false,false,false,false,false,false,false,false,false,true,false,false,false,true,false,false,true,false,false,true,true,true,true,true,true,true,false,false,true,true,false,false,false,false,false]
//[false,false,true,false,false,false,false,false,false,false,true,true,true,true,true,true,false,false,false,false,false,false,false,false,false,false,false,false,false,false,false,false,false,false,false,false,false,false,false,false,false,false,false,true,false,false,false,true,false,false,true,false,false,true,true,true,true,true,true,true,false,false,true,false,false,false,false,true,true]
//[false false true false false false false false false false true true true true true true false false false false false false false false false false false false false false false false false false false false false false false false false false false true false false false true false false true false false true true true true true true true false false true false false false false false false]
//17330550896613838894
//16214357719518126172
//13981971365326700729
//9517198656943849842

//[false,false,false,false,false,true,false,false,false,true,false,false,true,false,false,false,false,true,true]
//[false false false false false true false false false true true true false false false false false false false]
