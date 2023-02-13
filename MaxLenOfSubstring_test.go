package main

import (
	"testing"
)

// 字符串:最长不重复子串

func TestMaxLenOfSubstring(t *testing.T) {
	println(leon("abcabcdefg"))
	println(lengthOfLongestSubstring("abcabcdefg"))

}

func leon(s string) int {

	m := map[byte]int{}
	length := len(s)

	ret, left := 0, 0
	for i := 0; i < length; i++ {

		if m[s[i]] == 1 {

			ret = i - left
			delete(m, s[left])
			left++

		} else {
			m[s[i]] = 1
		}
	}
	ret = max(ret, length-left)

	return ret

}
func lengthOfLongestSubstring(s string) int {
	// 哈希集合，记录每个字符是否出现过
	m := map[byte]int{}
	n := len(s)
	// 右指针，初始值为 -1，相当于我们在字符串的左边界的左侧，还没有开始移动
	rk, ans := -1, 0
	for i := 0; i < n; i++ {
		if i != 0 {
			// 左指针向右移动一格，移除一个字符
			delete(m, s[i-1])
		}
		for rk+1 < n && m[s[rk+1]] == 0 {
			// 不断地移动右指针
			m[s[rk+1]]++
			rk++
		}
		// 第 i 到 rk 个字符是一个极长的无重复字符子串
		ans = max(ans, rk-i+1)
	}
	return ans
}

func max(x, y int) int {
	if x < y {
		return y
	}
	return x
}

// func lengthOfLongestSubstring(s string) int {
// 	if len(s) == 0 {
// 		return 0
// 	}
// 	var freq [256]int
// 	result, left, right := 0, 0, -1
// 	for left < len(s) {
// 		if right+1 < len(s) && freq[s[right+1]-'a'] == 0 {
// 			freq[s[right+1]-'a']++
// 			right++
// 		} else {
// 			freq[s[left]-'a']--
// 			left++
// 		}
// 		result = max(result, right-left+1)
// 	}
// 	return result
// }
// func max(a int, b int) int {
// 	if a > b {
// 		return a
// 	}
// 	return b
// }
