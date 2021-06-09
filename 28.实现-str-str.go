/*
 * @lc app=leetcode.cn id=28 lang=golang
 *
 * [28] 实现 strStr()
 */

// @lc code=start

func strStr(haystack string, needle string) int {
	if needle == "" {
		return 0
	}
	l1 := len(haystack)
	l2 := len(needle)
	max := l1 - l2
	for i := 0; i <= max; i++ {
		if haystack[i:i+len(needle)] == needle {
			return i
		}
	}
	return -1
}

// @lc code=end

