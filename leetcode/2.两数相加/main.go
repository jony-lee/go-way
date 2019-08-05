package main

import (
	"fmt"
	"strconv"
	"strings"
)

/*
给出两个 非空 的链表用来表示两个非负的整数。其中，它们各自的位数是按照 逆序 的方式存储的，并且它们的每个节点只能存储 一位 数字。

如果，我们将这两个数相加起来，则会返回一个新的链表来表示它们的和。

您可以假设除了数字 0 之外，这两个数都不会以 0 开头。

示例：

输入：(2 -> 4 -> 3) + (5 -> 6 -> 4)
输出：7 -> 0 -> 8
原因：342 + 465 = 807

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/add-two-numbers
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。


解题思路：
//最开始思路，实现链表的加减法，因为一些bug导致百倍，有些复杂
第二次打算先将链表转成整数进行操作，再写一个函数反向转换
利用int64进行操作依然不行，因为有超过了int64大小的数
最后还是回到链表计算中来
题目难点主要在于for循环跳出和if语句的一些边界判断
总体思路是先创建根节点，然后for循环，跳出循环条件是l1、l2均指向nil
创建两个待加变量t1、t2，如果l1（或l2）不是nil就将t1（或t2）赋值给l1.Val（或l2.Val）
然后构造新节点存储加和的值，当然这里需要注意进位问题
然后将新节点接在待返回节点的尾部
当所有数加完之后，判断一下进位
最终，根节点的Next节点返回即可
时间效率O()
空间效率O()
总结：主要是对单链表的一次应用，题目本身不难，按部就班地实现，考虑清楚边界条件。
*/
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
type ListNode struct {
	Val  int
	Next *ListNode
}

func (l *ListNode) String() string {
	s := []string{}
	for {
		s = append(s, strconv.Itoa(l.Val))
		if l.Next == nil {
			break
		}
		l = l.Next
	}
	return strings.Join(s, "->")
}

//思路1，成功-------------------------------------------------------

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	root := &ListNode{}
	node := root
	plus := 0 //进位符
	for l1 != nil || l2 != nil {
		var t1, t2 int //如果节点是nil，t1，t2就默认是0
		if l1 != nil {
			t1 = l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			t2 = l2.Val
			l2 = l2.Next
		}

		temp := t1 + t2 + plus
		plus = temp / 10
		tmpNode := &ListNode{
			Val:  temp % 10,
			Next: nil,
		}
		node.Next = tmpNode
		node = tmpNode
	}
	//最高位进1
	if plus > 0 {
		node.Next = &ListNode{plus, nil}
	}
	return root.Next
}

//思路2：-------------------------------------------------------

//
//func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
//	return IntToList(ListToInt(l1) + ListToInt(l2))
//}
//链表转整数
//func ListToInt(l1 *ListNode) int64{
//	var n,i int64 =0,1		//n是整数、i是当前位数
//	for {
//		n=i*int64(l1.Val) + n
//		i *= 10
//		if l1.Next == nil{
//			break
//		}
//		l1 = l1.Next
//	}
//	//n = ConverseInt(n)
//	return n
//}

//整数转链表
//func IntToList(n int64) *ListNode{
//	root := &ListNode{}
//	node := root
//	for {
//		node.Val = int(n % 10)
//		n = n / 10
//		if n == 0{
//			break
//		}
//		node.Next = &ListNode{}
//		node = node.Next
//	}
//	return root
//}
////整数反转
//func ConverseInt(n int) int {
//	m := 0
//	for n != 0{
//		m = m * 10 + n % 10
//		n = n/10
//	}
//	return m
//}

//他人优秀代码----------------------------------

//-------------------------------------------------------

func main() {
	l1 := &ListNode{
		Val: 2,
		Next: &ListNode{
			Val:  4,
			Next: nil,
		},
	}
	l2 := &ListNode{
		Val: 5,
		Next: &ListNode{
			Val: 6,
			Next: &ListNode{
				Val:  4,
				Next: nil,
			},
		},
	}
	fmt.Println(l1, l2)
	fmt.Println(addTwoNumbers(l1, l2))
}
