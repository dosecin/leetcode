/*
 * @lc app=leetcode.cn id=268 lang=golang
 *
 * [268] 缺失数字
 */

// @lc code=start
func missingNumber(nums []int) int {
	sum := 0
	for i, num := range nums {
		sum += i
		sum -= num
	}
	return sum + len(nums)
}

// @lc code=end
