package main

import (
	"fmt"
	"strconv"
	"strings"
)

/*
【题目】


【思路】
//TODO



*/

//【代码】
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
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	node := head
	root := &ListNode{
		Val:  0,
		Next: head,
	}
	delNode := root
	for node.Next != nil {
		node = node.Next
		if n > 1 {
			n--
			continue
		}
		delNode = delNode.Next
	}

	delNode.Next = delNode.Next.Next
	return root.Next
}
func isMatch(s string, p string) bool {
	if len(s) == 0 && (len(p) == 0 || p == "*") { //字符串为空，匹配为空或*可成功
		return true
	}
	if len(s) != 0 && len(p) == 0 { //字符串不为空，匹配为空则失败
		return false
	}

	j := 0
	for i := 0; i < len(s); i++ {
		if s[i] !=p[j]&&p[j]!='?'&&p[j]!='*'{
			return false
		}

		if s[i]==p[j]||p[j]=='?'{
			j++
			continue
		}else if p[j]=='*'{
			continue
		}

	}
	return true
}

//【主函数】
func main() {
	//a:= &ListNode{
	//	1,
	//	&ListNode{
	//		2,
	//		&ListNode{
	//			3,
	//			&ListNode{
	//				4,
	//				&ListNode{
	//					5,
	//					nil,
	//				},
	//			},
	//		},
	//	},
	//}
	//fmt.Println(removeNthFromEnd(a,2))

	s := "asd"
	p := "asd"
	fmt.Println(isMatch(s, p))
}

/*
【总结】：






*/
boolean isMatch(String str, String pattern) {
int s = 0, p = 0, match = 0, starIdx = -1;
//遍历整个字符串
while (s < str.length()){
// 一对一匹配，两指针同时后移。
if (p < pattern.length()  && (pattern.charAt(p) == '?' || str.charAt(s) == pattern.charAt(p))){
s++;
p++;
}
// 碰到 *，假设它匹配空串，并且用 startIdx 记录 * 的位置，记录当前字符串的位置，p 后移
else if (p < pattern.length() && pattern.charAt(p) == '*'){
starIdx = p;
match = s;
p++;
}
// 当前字符不匹配，并且也没有 *，回退
// p 回到 * 的下一个位置
// match 更新到下一个位置
// s 回到更新后的 match
// 这步代表用 * 匹配了一个字符
else if (starIdx != -1){
p = starIdx + 1;
match++;
s = match;
}
//字符不匹配，也没有 *，返回 false
else return false;
}

//将末尾多余的 * 直接匹配空串 例如 text = ab, pattern = a*******
while (p < pattern.length() && pattern.charAt(p) == '*')
p++;

return p == pattern.length();
}

作者：windliang
链接：https://leetcode-cn.com/problems/two-sum/solution/xiang-xi-tong-su-de-si-lu-fen-xi-duo-jie-fa-by-w-9/
来源：力扣（LeetCode）
著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。