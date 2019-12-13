/*
 * @lc app=leetcode.cn id=79 lang=golang
 *
 * [79] 单词搜索
 */

// @lc code=start
func exist(board [][]byte, word string) bool {
	if len(word) == 0 {
		return true
	}
	m, n := len(board), len(board[0])
	var dfs func(i, j int, word string) bool
	dfs = func(i, j int, word string) bool {
		if len(word) == 0 {
			return true
		}
		if i < 0 || i >= m || j < 0 || j >= n {
			return false
		}
		if board[i][j] != word[0] {
			return false
		}
		board[i][j] = '*'
		if dfs(i-1, j, word[1:]) {
			return true
		}
		if dfs(i, j-1, word[1:]) {
			return true
		}
		if dfs(i+1, j, word[1:]) {
			return true
		}
		if dfs(i, j+1, word[1:]) {
			return true
		}
		board[i][j] = word[0]
		return false
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if dfs(i, j, word) {
				return true
			}
		}
	}
	return false
}

// @lc code=end
