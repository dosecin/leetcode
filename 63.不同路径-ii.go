/*
 * @lc app=leetcode.cn id=63 lang=golang
 *
 * [63] 不同路径 II
 */

// @lc code=start
func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	if len(obstacleGrid) == 0 || len(obstacleGrid[0]) == 0 || obstacleGrid[0][0] != 0 {
		return 0
	}
	m, n := len(obstacleGrid), len(obstacleGrid[0])
	if obstacleGrid[m-1][n-1] != 0 {
		return 0
	}
	step := make([]int, m, m)
	step[0] = 1
	for i := 0; i < n; i++ {
		if obstacleGrid[0][i] != 0 {
			step[0] = 0
		}
		for j := 1; j < m; j++ {
			iAddUp := 0
			if obstacleGrid[j-1][i] == 0 {
				iAddUp = step[j-1]
			}
			iAddLeft := 0
			if i > 0 && obstacleGrid[j][i-1] == 0 {
				iAddLeft = step[j]
			}
			step[j] = iAddUp + iAddLeft
		}
	}
	return step[m-1]
}

// @lc code=end
