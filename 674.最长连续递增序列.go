/*
 * @lc app=leetcode.cn id=674 lang=golang
 *
 * [674] 最长连续递增序列
 */

// @lc code=start
func findLengthOfLCIS(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	max, cnt := 1, 1
	for i := 1; i < len(nums); i++ {
		if nums[i] <= nums[i-1] {
			cnt = 1
		} else {
			cnt++
		}
		if cnt > max {
			max = cnt
		}
	}
	return max
}

// @lc code=end
