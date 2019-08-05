package main

import "fmt"

/*
【题目】
有两种特殊字符。第一种字符可以用一比特0来表示。第二种字符可以用两比特(10 或 11)来表示。

现给一个由若干比特组成的字符串。问最后一个字符是否必定为一个一比特字符。给定的字符串总是由0结束。

示例 1:

输入:
bits = [1, 0, 0]
输出: True
解释:
唯一的编码方式是一个两比特字符和一个一比特字符。所以最后一个字符是一比特字符。
示例 2:

输入:
bits = [1, 1, 1, 0]
输出: False
解释:
唯一的编码方式是两比特字符和两比特字符。所以最后一个字符不是一比特字符。
注意:

1 <= len(bits) <= 1000.
bits[i] 总是0 或 1.

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/1-bit-and-2-bit-characters
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

【思路】
//TODO
其实是一道找规律题，去掉最右边第一个0，大致规律是数组从右往左数1的个数，如果遇到0，则结束for，当1是奇数个的时候不合题意
当1是偶数个的时候符合为true
另外给出的解法中有个关于异或的技巧，就是一个数不断异或另一个数的化，每两次异或就会回到自身值
由于go语言中的bool型变量不支持异或操作，所以用了整型，然后再作判断，否则的化，用布尔型异或看起来更加优雅。

*/

//【代码】
func isOneBitCharacter(bits []int) bool {
	ret := 1
	for i := len(bits) - 2; i >= 0; i-- {
		if bits[i] == 0 {
			break
		}
		ret = ret ^ 1
	}
	if ret == 1 {
		return true
	}
	return false
}

//【主函数】
func main() {
	a := []int{1, 0, 0}
	fmt.Println(isOneBitCharacter(a))
}

/*
【总结】：
题目较为简单，找到规律编写代码即可。





*/
