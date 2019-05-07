/*
 * @lc app=leetcode.cn id=414 lang=golang
 *
 * [414] 第三大的数
 */
func thirdMax(nums []int) int {
	max := [3]int{math.MinInt64, math.MinInt64, math.MinInt64}

	for _, n := range nums {
		if n > max[0] {
			max[0], max[1], max[2] = n, max[0], max[1]
		} else if n < max[0] && n > max[1] {
			max[1], max[2] = n, max[1]
		} else if n < max[1] && n > max[2] {
			max[2] = n
		}
	}
	if max[1] == math.MinInt64 || max[2] == math.MinInt64 {
		return max[0]
	}
	return max[2]
}

