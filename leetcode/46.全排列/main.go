package main

import "fmt"

/*
【题目】
给定一个没有重复数字的序列，返回其所有可能的全排列。

示例:

输入: [1,2,3]
输出:
[
  [1,2,3],
  [1,3,2],
  [2,1,3],
  [2,3,1],
  [3,1,2],
  [3,2,1]
]

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/permutations
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

【思路】
//TODO
从冒泡排序中找到一点灵感，如果要全排列只需要挨个让第一个元素冒泡到末尾就可以实现了【思路好像错了】
好像不仅要相邻交换，还要隔着交换
举个例子
[1,2,3]
1开始冒泡
[2,1,3]  [2,3,1]
然后2
[3,2,1] [3,1,2]
然后3
[1,3,2] [1,2,3]
有没有发现，当每个元素都冒过一遍泡之后，全排列就都产生了，并且数组元素顺序又回到了原来的状态。
*/

//【代码】
func permute(nums []int) [][]int {
	length:= len(nums)
	if length == 0{
		return [][]int{}
	}
	if length == 1{

		return [][]int{nums}
	}
	ret:= [][]int{}

	for i:=0;i<length;i++{
		for j:=0;j<length;j++{
			nums[i],nums[j] =nums[j],nums[i]
			temp:=make([]int, length)
			copy(temp,nums)
			ret = append(ret,temp)
		}
	}
	return ret
}
//【主函数】
func main() {
	a:=[]int{1,2,3,4}
	fmt.Println(permute(a))
}

/*
【总结】：






*/
