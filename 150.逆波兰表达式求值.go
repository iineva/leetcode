/*
 * @lc app=leetcode.cn id=150 lang=golang
 *
 * [150] 逆波兰表达式求值
 */

// @lc code=start
func evalRPN(tokens []string) int {
	stack := []int{}
	for _, v := range tokens {
		switch v {
		case "+":
			l := len(stack)
			if l < 2 {
				return -1
			}
			stack[l-2] = stack[l-2] + stack[l-1]
			stack = stack[:l-1]
		case "-":
			l := len(stack)
			if l < 2 {
				return -1
			}
			stack[l-2] = stack[l-2] - stack[l-1]
			stack = stack[:l-1]
		case "*":
			l := len(stack)
			if l < 2 {
				return -1
			}
			stack[l-2] = stack[l-2] * stack[l-1]
			stack = stack[:l-1]
		case "/":
			l := len(stack)
			if l < 2 {
				return -1
			}
			stack[l-2] = stack[l-2] / stack[l-1]
			stack = stack[:l-1]
		default:
			stack = append(stack, parseInt(v))
		}
	}
	return stack[0]
}

func parseInt(s string) int {
	if s == "" {
		return 0
	}
	b := 1
	if s[0] == '-' {
		b = -1
	}
	n := 0
	i := 0
	if b == -1 {
		i++
	}
	for ; i < len(s); i++ {
		if s[i] == ' ' {
			continue
		}
		if s[i] < '0' || s[i] > '9' {
			return n * b
		}
		n = n*10 + int(s[i]-'0')
	}
	return n * b
}

// @lc code=end

