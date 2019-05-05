import "strings"

/*
 * @lc app=leetcode.cn id=290 lang=golang
 *
 * [290] 单词模式
 */
func wordPattern(pattern string, str string) bool {
	strs := strings.Split(str, " ")
	if len(strs) != len(pattern) {
		return false
	}
	s2p := map[string]byte{}
	p2s := map[byte]string{}
	for i := range pattern {
		p := pattern[i]
		s := strs[i]
		b, ok := s2p[s]
		if ok {
			if b != p {
				return false
			}
		} else {
			_, ok = p2s[p]
			if ok {
				return false
			}
		}
		s2p[s] = p
		p2s[p] = s
	}
	return true
}
