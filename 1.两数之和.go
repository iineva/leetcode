/*
 * @lc app=leetcode.cn id=1 lang=golang
 *
 * [1] 两数之和
 */

// @lc code=start
func twoSum(nums []int, target int) []int {
	m := map[int]int{}
	for i, n := range nums {
		if v, ok := m[n]; ok {
			return []int{v, i}
		}
		m[target-n] = i
	}
	return []int{-1, -1}
}

// @lc code=end

