/*
 * @lc app=leetcode.cn id=541 lang=golang
 *
 * [541] 反转字符串 II
 */
func reverseStr(s string, k int) string {
	b, pos := []byte(s), 0
	for pos < len(b) {
		l := pos + k
		if l > len(b) {
			l = len(b)
		}
		for i, j := pos, l-1; i < j; i, j = i+1, j-1 {
			b[i], b[j] = b[j], b[i]
		}
		pos += 2 * k
	}
	return string(b)
}

