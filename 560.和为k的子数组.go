/*
 * @lc app=leetcode.cn id=560 lang=golang
 *
 * [560] 和为K的子数组
 */

// @lc code=start
func subarraySum(nums []int, k int) int {
	allSum := map[int]int{0: 1}
	ans := 0
	sum := 0
	for _, num := range nums {
		sum += num
		diff := sum - k
		if cnt, ok := allSum[diff]; ok {
			ans += cnt
		}
		if _, ok := allSum[sum]; ok {
			allSum[sum]++
		} else {
			allSum[sum] = 1
		}
	}
	return ans
}

// @lc code=end
