/*
 * @lc app=leetcode.cn id=209 lang=golang
 *
 * [209] 长度最小的子数组
 */

// @lc code=start
func minSubArrayLen(s int, nums []int) int {
	pre, cur := 0, 0
	sum := 0
	minLen := 0
	for pre < len(nums) && cur < len(nums) {
		if sum < s {
			sum += nums[cur]
			cur++
		} else {
			curLen := cur - pre
			if minLen == 0 || curLen < minLen {
				minLen = curLen
			}
			sum -= nums[pre]
			pre++
		}
	}
	if sum >= s && minLen == 0 {
		minLen = cur
	}
	for pre < len(nums) {
		sum -= nums[pre]
		pre++
		if sum >= s {
			curLen := cur - pre
			if minLen == 0 || curLen < minLen {
				minLen = curLen
			}
		}
	}
	return minLen
}

// @lc code=end
