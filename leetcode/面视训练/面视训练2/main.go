package main

//给定一个排序数组和一个目标值，在数组中找到目标值，并返回其索引。如果目标值不存在于数组中，返回它将会被按顺序插入的位置。
//
//你可以假设数组中无重复元素。

//func searchInsert(nums []int, target int) int {
//	//二分查找排序
//	length := len(nums)
//	if target < nums[0]{
//		return 0
//	}
//	if target > nums[length-1]{
//		return length
//	}
//
//	low,high := 0,length-1
//	k := length / 2
//	for high-low>1{
//		if target >= nums[k]{
//			low = k
//			k = (high + low + 1)/2
//		}else{
//			high = k
//			k = (high + low + 1)/2
//		}
//	}
//	if nums[low]<target{
//		low = low + 1
//	}
//	return low
//}

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
//func levelOrder(root *TreeNode) [][]int {
//	if root == nil{
//		return [][]int{}
//	}
//
//}

func maxChunksToSorted(arr []int) int {

}

func main() {
	//nums:= []int{1,3,5,6}
	//target:= 2
	//fmt.Println(searchInsert(nums,target))

}
