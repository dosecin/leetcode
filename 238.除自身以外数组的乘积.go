/*
 * @lc app=leetcode.cn id=238 lang=golang
 *
 * [238] 除自身以外数组的乘积
 */

// @lc code=start
func productExceptSelf(nums []int) []int {
	ans := make([]int, len(nums))
	ans[0] = 1
	for i := 1; i < len(nums); i++ {
		ans[i] = ans[i-1] * nums[i-1]
	}
	mul := 1
	for i := len(nums) - 2; i >= 0; i-- {
		mul = mul * nums[i+1]
		ans[i] *= mul
	}
	return ans
}

// @lc code=end
