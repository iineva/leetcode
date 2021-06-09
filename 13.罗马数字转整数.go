/*
 * @lc app=leetcode.cn id=13 lang=golang
 *
 * [13] 罗马数字转整数
 */

// @lc code=start
// 数据量少 switch 比map快！
func get(k rune) int {
	switch k {
	case 'I':
		return 1
	case 'V':
		return 5
	case 'X':
		return 10
	case 'L':
		return 50
	case 'C':
		return 100
	case 'D':
		return 500
	case 'M':
		return 1000
	}
	return 0
}

func romanToInt(s string) int {
	n := 0
	for i, v := range s {
		if i < (len(s)-1) && get(v) < get(rune(s[i+1])) {
			n -= get(v)
		} else {
			n += get(v)
		}
	}
	return n
}

// @lc code=end

