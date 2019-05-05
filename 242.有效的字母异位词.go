/*
 * @lc app=leetcode.cn id=242 lang=golang
 *
 * [242] 有效的字母异位词
 */
func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	srune := map[rune]int{}
	for _, v := range s {
		srune[v]++
	}
	for _, v := range t {
		srune[v]--
		if srune[v] < 0 {
			return false
		}
		if srune[v] == 0 {
			delete(srune, v)
		}
	}
	return len(srune) == 0
}
