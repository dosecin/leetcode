/*
 * @lc app=leetcode.cn id=485 lang=golang
 *
 * [485] 最大连续1的个数
 */

// @lc code=start
func findMaxConsecutiveOnes(nums []int) int {
	max, cnt := 0, 0
	for _, num := range nums {
		if num == 0 {
			cnt = 0
		} else {
			cnt++
			if cnt > max {
				max = cnt
			}
		}
	}
	return max
}

// @lc code=end
