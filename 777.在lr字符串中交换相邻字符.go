import "strings"

/*
 * @lc app=leetcode.cn id=777 lang=golang
 *
 * [777] 在LR字符串中交换相邻字符
 */

// @lc code=start
func canTransform(start string, end string) bool {
	s := strings.Replace(start, "X", "", -1)
	e := strings.Replace(end, "X", "", -1)
	if s != e {
		return false
	}

	tL, tR := 0, 0
	for i, b := range []byte(start) {
		if b == 'L' {
			for tL < len(end) {
				if end[tL] == 'L' {
					break
				}
				tL++
			}
			if i < tL {
				return false
			}
			tL++
		} else if b == 'R' {
			for tR < len(end) {
				if end[tR] == 'R' {
					break
				}
				tR++
			}
			if i > tR {
				return false
			}
			tR++
		}
	}
	return true
}

// @lc code=end
