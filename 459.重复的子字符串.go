/*
 * @lc app=leetcode.cn id=459 lang=golang
 *
 * [459] 重复的子字符串
 */
func repeatedSubstringPattern(s string) bool {
	if len(s) <= 1 {
		return false
	}
	pos := 1
	for pos < len(s) {
		if s[pos] != s[0] {
			break
		}
		pos++
	}
	if pos == len(s) {
		return true
	}
	end := len(s) / 2
	for pos < end {
		for pos < end {
			if s[pos] == s[0] {
				break
			}
			pos++
		}
		if len(s)%pos != 0 {
			pos++
			continue
		}
		i, j := 0, pos
		for j < len(s) {
			if s[i] != s[j] {
				break
			}
			j++
			i = (i + 1) % pos
		}
		if j == len(s) {
			return true
		}
		pos++
	}
	return false
}

