/*
 * @lc app=leetcode.cn id=495 lang=golang
 *
 * [495] 提莫攻击
 */

// @lc code=start
func findPoisonedDuration(timeSeries []int, duration int) int {
	if len(timeSeries) == 0 {
		return 0
	}
	sum := duration
	for i := 1; i < len(timeSeries); i++ {
		diff := timeSeries[i] - timeSeries[i-1]
		if diff > duration {
			diff = duration
		}
		sum += diff
	}
	return sum
}

// @lc code=end
