/*
 * @lc app=leetcode.cn id=41 lang=golang
 *
 * [41] 缺失的第一个正数
 */

// @lc code=start
func firstMissingPositive(nums []int) int {
	haveOne := false
	for i, n := range nums {
		if n == 1 {
			haveOne = true
		}
		if n <= 0 || n > len(nums) {
			nums[i] = 1
		}
	}
	if !haveOne {
		return 1
	}
	if len(nums) == 1 {
		return 2
	}
	for _, n := range nums {
		if n < 0 {
			n = -n
		}
		n = n - 1
		if nums[n] > 0 {
			nums[n] = -nums[n]
		}
	}
	ret := 0
	for ret < len(nums) {
		if nums[ret] > 0 {
			return ret + 1
		}
		ret++
	}
	return ret + 1
}

// @lc code=end
