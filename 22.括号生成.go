/*
 * @lc app=leetcode.cn id=22 lang=golang
 *
 * [22] 括号生成
 */

// @lc code=start

func generateParenthesis(n int) []string {

	cache := map[string]bool{}
	g([]byte{'('}, n, 1, 0, cache)

	s := []string{}
	for k, _ := range cache {
		s = append(s, k)
	}
	return s
}

func g(s []byte, n, l, r int, cache map[string]bool) {
	if len(s) == n*2 {
		if isVaid(s) {
			cache[string(s)] = true
		}
	} else {
		if l < n {
			g(append(s, '('), n, l+1, r, cache)
		}
		if r < n && r < l {
			g(append(s, ')'), n, l, r+1, cache)
		}
	}
}

func isVaid(s []byte) bool {
	l := 0
	for _, v := range s {
		if v == '(' {
			l++
		}
		if v == ')' {
			l--
		}
		if l < 0 {
			return false
		}
	}
	return l == 0
}

// @lc code=end

