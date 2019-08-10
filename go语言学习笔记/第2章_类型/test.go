package main

import "fmt"

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
