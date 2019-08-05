package main

import "fmt"

/*
【题目1】
给定一个按非递减顺序排序的整数数组 A，返回每个数字的平方组成的新数组，要求也按非递减顺序排序。



示例 1：

输入：[-4,-1,0,3,10]
输出：[0,1,9,16,100]
示例 2：

输入：[-7,-3,2,3,11]
输出：[4,9,9,49,121]

【题目2】
给出一些不同颜色的盒子，盒子的颜色由数字表示，即不同的数字表示不同的颜色。
你将经过若干轮操作去去掉盒子，直到所有的盒子都去掉为止。每一轮你可以移除具有相同颜色的连续 k 个盒子（k >= 1），这样一轮之后你将得到 k*k 个积分。
当你将所有盒子都去掉之后，求你能获得的最大积分和。

示例 1：
输入:

[1, 3, 2, 2, 2, 3, 4, 3, 1]
输出:

23
解释:

[1, 3, 2, 2, 2, 3, 4, 3, 1]
----> [1, 3, 3, 4, 3, 1] (3*3=9 分)
----> [1, 3, 3, 3, 1] (1*1=1 分)
----> [1, 1] (3*3=9 分)
----> [] (2*2=4 分)

【思路1】
//TODO
先找到最接近0的位置，然后利用两个指针分别向左右滑动，同时比较大小，若元素更小则求平方并存入返回值数组。
【思路2】

*/

//【代码1】完成
func sortedSquares(A []int) []int {
	k := 0

	ret := make([]int, len(A))
	for i, v := range A {
		if v >= 0 {
			k = i
			break
		}
	}
	neg, pos := k-1, k
	for i := 0; i < len(A); i++ {
		if neg < 0 {
			ret[i] = A[pos] * A[pos]
			pos += 1
			continue
		}
		if pos >= len(A) {
			ret[i] = A[neg] * A[neg]
			neg -= 1
			continue
		}
		if A[pos] < (-A[neg]) {
			ret[i] = A[pos] * A[pos]
			pos += 1
		} else {
			ret[i] = A[neg] * A[neg]
			neg -= 1
		}
	}
	return ret
}

//【代码2】
func removeBoxes(boxes []int) int {

}

//【代码3】
func totalNQueens(n int) int {
	if n == 0 {
		return 0
	}

	if n%6 != 2 && n%6 != 3 {
		return 2
	} else {
		return n / 2
	}

}

//【主函数】
func main() {
	//【解1】
	//a:= []int{-4,-1,0,3,10}
	//fmt.Println(sortedSquares(a))
	//【解2】
	//box:=[]int{1, 3, 2, 2, 2, 3, 4, 3, 1}
	//fmt.Println(removeBoxes(box))
	//【解3】
	n := 5
	fmt.Println(totalNQueens(5))
}

/*
【总结】：






*/
