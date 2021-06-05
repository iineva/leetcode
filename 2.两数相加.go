/*
 * @lc app=leetcode.cn id=2 lang=golang
 *
 * [2] 两数相加
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	n := l1.Val + l2.Val
	b := false
	if n >= 10 {
		b = true
		n = n - (n / 10 * 10)
	}
	node := &ListNode{Val: n}
	l := node
	for l1.Next != nil || l2.Next != nil || b {
		l1 = l1.Next
		l2 = l2.Next
		if l1 == nil {
			l1 = &ListNode{}
		}
		if l2 == nil {
			l2 = &ListNode{}
		}
		n = l1.Val + l2.Val
		if b {
			n += 1
		}
		if n >= 10 {
			b = true
			n = n - (n / 10 * 10)
		} else {
			b = false
		}
		l.Next = &ListNode{Val: n}
		l = l.Next
	}
	return node
}

// @lc code=end

