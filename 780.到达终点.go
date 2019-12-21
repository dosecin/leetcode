/*
 * @lc app=leetcode.cn id=780 lang=golang
 *
 * [780] 到达终点
 */

// @lc code=start
func reachingPoints(sx int, sy int, tx int, ty int) bool {
	for tx >= sx && ty >= sy {
		if tx > ty {
			tx = tx - ty
		} else {
			ty = ty - tx
		}
	}
	return sx == tx && sy == ty
}

// @lc code=end
