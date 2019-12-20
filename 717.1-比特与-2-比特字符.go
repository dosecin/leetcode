/*
 * @lc app=leetcode.cn id=717 lang=golang
 *
 * [717] 1比特与2比特字符
 */

// @lc code=start
func isOneBitCharacter(bits []int) bool {
	cur := 0
	for cur < len(bits) {
		if cur+1 == len(bits) {
			return true
		}
		if bits[cur] == 0 {
			cur++
		} else {
			cur += 2
		}
	}
	return false
}

// @lc code=end
