import "math"

/*
 * @lc app=leetcode.cn id=152 lang=golang
 *
 * [152] 乘积最大子序列
 */

// @lc code=start
func maxProduct(nums []int) int {
	max := func(ans int, nums ...int) int {
		for _, num := range nums {
			if num > ans {
				ans = num
			}
		}
		return ans
	}
	min := func(ans int, nums ...int) int {
		for _, num := range nums {
			if num < ans {
				ans = num
			}
		}
		return ans
	}
	ans := math.MinInt64
	dpMax, dpMin := 1, 1
	for _, num := range nums {
		dpMax, dpMin = max(dpMax*num, num, num*dpMin), min(num, dpMin*num, dpMax*num)
		if dpMax > ans {
			ans = dpMax
		}
	}
	return ans
}

// @lc code=end
