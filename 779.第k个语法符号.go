/*
 * @lc app=leetcode.cn id=779 lang=golang
 *
 * [779] 第K个语法符号
 */

// @lc code=start
func kthGrammar(N int, K int) int {
	if K == 1 {
		return 0
	}
	if K == 2 {
		return 1
	}
	g := kthGrammar(N, (K+1)/2)
	if g == 0 {
		if K%2 == 0 {
			return 1
		}
		return 0
	}
	if K%2 == 0 {
		return 0
	}
	return 1
}

// @lc code=end
