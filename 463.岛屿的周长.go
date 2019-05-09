/*
 * @lc app=leetcode.cn id=463 lang=golang
 *
 * [463] 岛屿的周长
 */
func islandPerimeter(grid [][]int) int {
	res := 0
	for i := range grid {
		for j := range grid[i] {
			if grid[i][j] == 0 {
				continue
			}
			res += 4
			if i > 0 && grid[i-1][j] == 1 {
				res--
			}
			if i < len(grid)-1 && grid[i+1][j] == 1 {
				res--
			}
			if j > 0 && grid[i][j-1] == 1 {
				res--
			}
			if j < len(grid[i])-1 && grid[i][j+1] == 1 {
				res--
			}
		}
	}
	return res
}

