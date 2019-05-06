/*
 * @lc app=leetcode.cn id=387 lang=golang
 *
 * [387] 字符串中的第一个唯一字符
 */
func firstUniqChar(s string) int {
	m := map[rune]int{}
	for _, r := range s {
		m[r]++
	}
	for i, r := range s {
		if m[r] == 1 {
			return i
		}
	}
	return -1
}

