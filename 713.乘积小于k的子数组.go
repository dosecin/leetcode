/*
 * @lc app=leetcode.cn id=713 lang=golang
 *
 * [713] 乘积小于K的子数组
 */

// @lc code=start
func numSubarrayProductLessThanK(nums []int, k int) int {
	ans := 0
	mul := 1
	pre, cur := 0, 0
	for cur < len(nums) {
		mul *= nums[cur]
		if mul < k {
			ans += cur - pre + 1
		} else {
			for mul >= k && pre <= cur {
				mul /= nums[pre]
				pre++
			}
			if mul < k {
				ans += cur - pre + 1
			}
		}
		cur++
	}
	return ans
}

// @lc code=end
