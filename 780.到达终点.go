/*
 * @lc app=leetcode.cn id=780 lang=golang
 *
 * [780] 到达终点
 */

// @lc code=start
func reachingPoints(sx int, sy int, tx int, ty int) bool {
	for tx >= sx && ty >= sy {
		if tx == ty {
			break
		}
		if tx > ty {
			if ty > sy {
				tx %= ty
			} else {
				return (tx-sx)%ty == 0
			}
		} else {
			if tx > sx {
				ty %= tx
			} else {
				return (ty-sy)%tx == 0
			}
		}
	}
	return sx == tx && sy == ty
}

// @lc code=end
