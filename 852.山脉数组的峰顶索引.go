/*
 * @lc app=leetcode.cn id=852 lang=golang
 *
 * [852] 山脉数组的峰顶索引
 */

// @lc code=start
func peakIndexInMountainArray(A []int) int {
	for i, num := range A[1:] {
		if A[i] > num {
			return i
		}
	}
	return len(A) - 1
}

// @lc code=end
