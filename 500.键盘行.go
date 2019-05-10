/*
 * @lc app=leetcode.cn id=500 lang=golang
 *
 * [500] 键盘行
 */

var m = map[byte]int{}

func init() {
	kb := []byte{'q', 'w', 'e', 'r', 't', 'y', 'u', 'i', 'o', 'p'}
	for _, c := range kb {
		m[c] = 1
		m['A'+c-'a'] = 1
	}
	kb = []byte{'a', 's', 'd', 'f', 'g', 'h', 'j', 'k', 'l'}
	for _, c := range kb {
		m[c] = 2
		m['A'+c-'a'] = 2
	}
	kb = []byte{'z', 'x', 'c', 'v', 'b', 'n', 'm'}
	for _, c := range kb {
		m[c] = 4
		m['A'+c-'a'] = 4
	}
}

func findWords(words []string) []string {
	res := []string{}
	for i := range words {
		h := 0
		for j := range words[i] {
			k := m[words[i][j]]
			h |= k
		}
		if h == 1 || h == 2 || h == 4 {
			res = append(res, words[i])
		}
	}
	return res
}

