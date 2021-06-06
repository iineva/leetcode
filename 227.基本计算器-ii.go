/*
 * @lc app=leetcode.cn id=227 lang=golang
 *
 * [227] 基本计算器 II
 */

// @lc code=start

var (
	m = map[string]int{
		"+": 0,
		"-": 0,
		"*": 1,
		"/": 1,
		"(": 2,
		")": 2,
	}
)

type parser struct {
	s string
	i int
}

func (p *parser) parse() []string {
	return p.p(0)
}

func (p *parser) p(j int) []string {
	stack := []string{}
	b := false
	opt := []string{}
	n := ""
	str := p.s + " "
	for p.i = j; p.i < len(str); p.i++ {
		v := str[p.i]
		if v >= '0' && v <= '9' {
			b = true
			n = n + string(v)
			continue
		}
		if b {
			stack = append(stack, n)
			b = false
			n = ""
			if len(opt) == 2 {
				if m[opt[0]] >= m[opt[1]] {
					stack = append(stack, opt[0])
					opt = []string{opt[1]}
				} else {
					stack = append(stack, opt[1])
					opt = []string{opt[0]}
				}
			}
		}
		if v == ' ' {
			continue
		}
		if v == '(' {
			stack = append(stack, p.p(p.i+1)...)
			continue
		}
		if v == ')' {
			break
		}
		if len(opt) > 0 && m[opt[0]] >= m[string(v)] {
			stack = append(stack, opt[0])
			opt = opt[1:]
		}
		opt = append(opt, string(v))
	}

	if len(opt) == 2 {
		if m[opt[0]] >= m[opt[1]] {
			stack = append(stack, opt[0])
			opt = []string{opt[1]}
		} else {
			stack = append(stack, opt[1])
			opt = []string{opt[0]}
		}
	}
	if len(opt) == 1 {
		stack = append(stack, opt[0])
	}
	// log.Print(stack)
	return stack
}

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

func calculate(s string) int {

	p := &parser{s: s}

	return evalRPN(p.parse())
}

// @lc code=end

