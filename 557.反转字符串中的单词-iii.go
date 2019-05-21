/*
 * @lc app=leetcode.cn id=557 lang=golang
 *
 * [557] 反转字符串中的单词 III
 */
func reverseWords(s string) string {
    reverse := func(b []byte) {
		for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
			b[i], b[j] = b[j], b[i]
		}
	}
	b := []byte(s)
	start, end := 0, 0
	for end < len(b) {
		if b[end] == ' ' {
			reverse(b[start:end])
			end++
			start = end
			continue
		}
		end++
	}
	reverse(b[start:end])
	return string(b)
}

