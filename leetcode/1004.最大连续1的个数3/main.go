package main

import "fmt"

/*
【题目】
给定一个由若干 0 和 1 组成的数组 A，我们最多可以将 K 个值从 0 变成 1 。

返回仅包含 1 的最长（连续）子数组的长度。

 

示例 1：

输入：A = [1,1,1,0,0,0,1,1,1,1,0], K = 2
输出：6
解释：
[1,1,1,0,0,1,1,1,1,1,1]
粗体数字从 0 翻转到 1，最长的子数组长度为 6。
示例 2：

输入：A = [0,0,1,1,0,0,1,1,1,0,1,1,0,0,0,1,1,1,1], K = 3
输出：10
解释：
[0,0,1,1,1,1,1,1,1,1,1,1,0,0,0,1,1,1,1]
粗体数字从 0 翻转到 1，最长的子数组长度为 10。
 

提示：

1 <= A.length <= 20000
0 <= K <= A.length
A[i] 为 0 或 1 

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/max-consecutive-ones-iii
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

【思路】
//TODO



*/

//【代码】
func longestOnes(A []int, K int) int {
	start,maxStart,max := 0,0,0
	n:= K
	for i,v:=range A{
		if v == 0 {
			if n == 0{
				length := i-K - maxStart
				if max < length{
					max = length
					maxStart = start
				}
				start = i + 1
			}
			n--
		}
	}
}
//【主函数】
func main() {
	A:=[]int{1,1,1,0,0,0,1,1,1,1,0}
	k:=2
	fmt.Println(longestOnes(A,k))
}

/*
【总结】：






*/
