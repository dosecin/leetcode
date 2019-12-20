/*
 * @lc app=leetcode.cn id=747 lang=golang
 *
 * [747] 至少是其他数字两倍的最大数
 */

// @lc code=start
func dominantIndex(nums []int) int {
	preMax, curMax := 0, nums[0]
	ans := 0
	for i, num := range nums[1:] {
		if num > curMax {
			preMax, curMax = curMax, num
			ans = i + 1
		} else if num > preMax {
			preMax = num
		}
	}
	if curMax >= 2*preMax {
		return ans
	}
	return -1
}

// @lc code=end
