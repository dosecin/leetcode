import "math"

/*
 * @lc app=leetcode.cn id=628 lang=golang
 *
 * [628] 三个数的最大乘积
 */

// @lc code=start
func maximumProduct(nums []int) int {
	max0, max1, max2 := math.MinInt64, math.MinInt64, math.MinInt64
	min0, min1 := math.MaxInt64, math.MaxInt64
	for _, num := range nums {
		if num > max0 {
			max0, max1, max2 = num, max0, max1
		} else if num > max1 {
			max1, max2 = num, max1
		} else if num > max2 {
			max2 = num
		}
		if num < min0 {
			min0, min1 = num, min0
		} else if num < min1 {
			min1 = num
		}
	}
	lef := min0 * min1 * max0
	rig := max0 * max1 * max2
	if lef > rig {
		return lef
	}
	return rig
}

// @lc code=end
