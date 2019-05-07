/*
 * @lc app=leetcode.cn id=434 lang=golang
 *
 * [434] 字符串中的单词数
 */
func countSegments(s string) int {
	var cnt, st, pos int
outer:
	for pos < len(s) {
		c := s[pos]
		if c == ' ' {
			if st != pos {
				cnt++
			}
			for c == ' ' {
				pos++
				st = pos
				if pos >= len(s) {
					break outer
				}
				c = s[pos]
			}
		}
		pos++
	}
	if st != pos {
		cnt++
	}
	return cnt
}

