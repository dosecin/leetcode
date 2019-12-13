/*
 * @lc app=leetcode.cn id=66 lang=golang
 *
 * [66] 加一
 */

// @lc code=start
func plusOne(digits []int) []int {
	bCarry := false
	for i := len(digits) - 1; i >= 0; i-- {
		digits[i]++
		if digits[i] < 10 {
			bCarry = false
			break
		}
		bCarry = true
		digits[i] -= 10
	}
	if bCarry {
		return append([]int{1}, digits...)
	}
	return digits
}

// @lc code=end
