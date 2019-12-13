/*
 * @lc app=leetcode.cn id=64 lang=golang
 *
 * [64] 最小路径和
 */

// @lc code=start
func minPathSum(grid [][]int) int {
	if len(grid) == 0 || len(grid[0]) == 0 {
		return 0
	}
	m, n := len(grid), len(grid[0])
	sum := make([]int, m, m)
	sum[0] = grid[0][0]
	for i := 1; i < m; i++ {
		sum[i] = sum[i-1] + grid[i][0]
	}
	for j := 1; j < n; j++ {
		for i := 0; i < m; i++ {
			if i == 0 {
				sum[i] += grid[i][j]
			} else {
				if sum[i] < sum[i-1] {
					sum[i] += grid[i][j]
				} else {
					sum[i] = sum[i-1] + grid[i][j]
				}
			}
		}
	}
	return sum[m-1]
}

// @lc code=end
