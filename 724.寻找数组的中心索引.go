/*
 * @lc app=leetcode.cn id=724 lang=golang
 *
 * [724] 寻找数组的中心索引
 */

// @lc code=start
func pivotIndex(nums []int) int {
	sums := make([]int, len(nums))
	sum := 0
	for i, num := range nums {
		sum += num
		sums[i] = sum
	}

	for i := range sums {
		lef := 0
		if i > 0 {
			lef = sums[i-1]
		}
		if lef == sum-sums[i] {
			return i
		}
	}
	return -1
}

// @lc code=end
