package main

import "fmt"

/*
题目模板


解题思路：
//TODO

时间效率O()
空间效率O()
*/

//思路1-------------------------------------------

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

//他人优秀代码----------------------------------

//主函数-------------------------------------------

func main() {
	nums := []int{2, 4, 7, 15}
	target := 9
	fmt.Println(twoSum(nums, target))
}
