package main

import "fmt"

/*
题名描述：
给定一个整数数组 nums 和一个目标值 target，请你在该数组中找出和为目标值的那 两个 整数，并返回他们的数组下标。

你可以假设每种输入只会对应一个答案。但是，你不能重复利用这个数组中同样的元素。

示例:

给定 nums = [2, 7, 11, 15], target = 9

因为 nums[0] + nums[1] = 2 + 7 = 9
所以返回 [0, 1]

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/two-sum
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。


解题思路：
该题本质上是一个查找问题，查找满足特定运算的两个数的位置

首先会想到进行暴力计算，即进行两层for循环，第二层for循环进行条件判断，如果满足题意则结束程序并返回
当然这样的算法的性能不够好，对于查找问题，首先想到的效率为O(1)的hash算法，那么思考是否能通过字典来完成查找问题
首先利用for循环依次将数放入字典中待查，每个数放入字典内之前需要先确认字典中没有我想要的匹配对象，翻译成代码m[target-n]不存在
既然不存在，就可以先令m[target-n]=i，来记住当前元素在数组中的位置。
一旦发现m[target-n]存在，说明当前要查找的数与之前放入字典中的数能够构成匹配关系，则将当前状态下的i位置和字典值位置返回即可。

时间效率O(n)
空间效率O(1)
总结：查找满足特定运算的两个数的位置，首先去思考hash算法能不能实现想要的结果，如果能，一般考虑通过字典怎么实现。
*/
func twoSum(nums []int, target int) []int {
	m := make(map[int]int)

	for i, n := range nums {
		if _, ok := m[n]; ok == false {
			m[target-n] = i
			continue
		}
		return []int{m[n], i}
	}
	return []int{}
}

func main() {
	nums := []int{2, 4, 7, 15}
	target := 9
	fmt.Println(twoSum(nums, target))
}
