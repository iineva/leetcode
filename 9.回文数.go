/*
 * @lc app=leetcode.cn id=9 lang=golang
 *
 * [9] 回文数
 */

// @lc code=start
func isPalindrome(x int) bool {
	// 忽略负数
	if x < 0 {
		return false
	} else if x == 0 {
		return true
	}
	b := x
	n := 0 // 将x反转为n
	for b > 0 {
		n = n*10 + b%10
		b = b / 10
	}
	// 比对反转后结果
	return x == n
}

// @lc code=end

