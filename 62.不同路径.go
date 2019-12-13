/*
 * @lc app=leetcode.cn id=62 lang=golang
 *
 * [62] 不同路径
 */

// @lc code=start
func uniquePaths(m int, n int) int {
	step := make([]int, m, m)
	for i := 0; i < m; i++ {
		step[i] = 1
	}
	for i := 1; i < n; i++ {
		for j := 1; j < m; j++ {
			step[j] = step[j-1] + step[j]
		}
	}
	return step[m-1]
}

// @lc code=end
