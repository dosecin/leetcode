/*
 * @lc app=leetcode.cn id=383 lang=golang
 *
 * [383] 赎金信
 */
func canConstruct(ransomNote string, magazine string) bool {
	m := map[rune]int{}
	for _, r := range magazine {
		m[r]++
	}
	for _, r := range ransomNote {
		_, ok := m[r]
		if !ok {
			return false
		}
		m[r]--
		if m[r] < 0 {
			return false
		}
	}
	return true
}

