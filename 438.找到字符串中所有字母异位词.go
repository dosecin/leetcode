/*
 * @lc app=leetcode.cn id=438 lang=golang
 *
 * [438] 找到字符串中所有字母异位词
 */
func findAnagrams(s string, p string) []int {
	m := map[byte]int{}
	for i := range p {
		m[p[i]]++
	}
	res, pos, cnt := []int{}, 0, 0
	for pos < len(s) {
		st := pos - len(p)
		if st >= 0 {
			c := s[st]
			_, ok := m[c]
			if ok {
				m[c]++
				if m[c] > 0 {
					cnt--
				}
			}
		}
		c := s[pos]
		pos++
		_, ok := m[c]
		if !ok {
			continue
		}
		m[c]--
		if m[c] >= 0 {
			cnt++
		}
		if cnt == len(p) {
			res = append(res, pos-cnt)
		}

	}
	return res
}

