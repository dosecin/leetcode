/*
 * @lc app=leetcode.cn id=289 lang=golang
 *
 * [289] 生命游戏
 */

// @lc code=start
func gameOfLife(board [][]int) {
	m, n := len(board), len(board[0])
	isLive := func(i, j int) bool {
		if i < 0 || i >= m {
			return false
		}
		if j < 0 || j >= n {
			return false
		}
		return board[i][j] > 0
	}
	neighbors := [][]int{
		{-1, -1}, {-1, 0}, {-1, 1}, {0, -1},
		{0, 1}, {1, -1}, {1, 0}, {1, 1},
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			liveNum := 0
			for _, neighbor := range neighbors {
				if isLive(i+neighbor[0], j+neighbor[1]) {
					liveNum++
				}
			}
			if board[i][j] == 1 {
				if liveNum < 2 || liveNum > 3 {
					board[i][j] = 2 // 2标记要从1->0
				}
			} else {
				if liveNum == 3 {
					board[i][j] = -1 // -1标记要从0->1
				}
			}
		}
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if board[i][j] == 2 {
				board[i][j] = 0
			} else if board[i][j] == -1 {
				board[i][j] = 1
			}
		}
	}
}

// @lc code=end
