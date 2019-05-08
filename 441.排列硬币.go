/*
 * @lc app=leetcode.cn id=441 lang=golang
 *
 * [441] 排列硬币
 */
func arrangeCoins(n int) int {
	line := 1
	for n >= line {
		n -= line
		line++
	}
	return line-1
}

