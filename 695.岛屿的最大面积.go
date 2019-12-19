/*
 * @lc app=leetcode.cn id=695 lang=golang
 *
 * [695] 岛屿的最大面积
 */

// @lc code=start
func maxAreaOfIsland(grid [][]int) int {
	if len(grid) == 0 || len(grid[0]) == 0 {
		return 0
	}
	m, n := len(grid), len(grid[0])
	directs := [][]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}
	ans := 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] != 1 {
				continue
			}
			grid[i][j] = -1
			land := 0
			stack := [][]int{{i, j}}
			for len(stack) > 0 {
				land++
				data := stack[len(stack)-1]
				stack = stack[:len(stack)-1]
				for _, d := range directs {
					ii, jj := data[0]+d[0], data[1]+d[1]
					if ii < 0 || ii >= m {
						continue
					}
					if jj < 0 || jj >= n {
						continue
					}
					if grid[ii][jj] != 1 {
						continue
					}
					grid[ii][jj] = -1
					stack = append(stack, []int{ii, jj})
				}
			}
			if land > ans {
				ans = land
			}
		}
	}
	return ans
}

// @lc code=end
