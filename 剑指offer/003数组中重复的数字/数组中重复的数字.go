package main

func longestPalindrome(s string) string {
	if len(s) < 1 {
		return ""
	}
	st, e := 0, 0 // 记录切片点
	for i := 0; i < len(s); i ++ {
		L1, R1 := expandAroundCenter(s, i, i)
		L2, R2 := expandAroundCenter(s, i, i + 1)
		L, R := 0, 0
		if R1 - L1 > R2 - L2 {
			L, R = L1, R1
		} else {
			L, R = L2, R2
		}
		if R - L - 1 > e - st {
			st = L
			e = R + 1
		}

	}
	return s[st: e]
}

func expandAroundCenter(s string, left int, right int)(int, int) {
	L, R := left, right
	for L >= 0 && R < len(s) && s[L] == s[R] { // 注意 L是>=
		L --
		R ++
	}
	return L, R
}