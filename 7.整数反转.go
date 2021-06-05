/*
 * @lc app=leetcode.cn id=7 lang=golang
 *
 * [7] 整数反转
 */

// @lc code=start
func reverse(x int) int {
	n := 0
	b := x < 0
	// 负数情况统一为正数
	if b {
		x = -x
	}

	for x > 0 {
		n = (n * 10) + (x % 10)
		x = x / 10
		// 判断超过 2^31 − 1 就返回0
		if n > (0xFFFFFFFF/2 - 1) {
			return 0
		}
	}

	// 负数情况反转
	if b {
		n = -n
	}

	return n
}

// @lc code=end

