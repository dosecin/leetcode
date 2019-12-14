/*
 * @lc app=leetcode.cn id=126 lang=golang
 *
 * [126] 单词接龙 II
 */

// @lc code=start
func findLadders(beginWord string, endWord string, wordList []string) [][]string {
	if beginWord == endWord {
		return [][]string{{beginWord}}
	}
	ans := [][]string{}
	canTransform := func(a, b string) bool {
		diff := 0
		for i := range a {
			if a[i] != b[i] {
				diff++
			}
		}
		return diff == 1
	}

	for i, str := range wordList {
		if !canTransform(beginWord, str) {
			continue
		}
		wl := append([]string{}, wordList[:i]...)
		wl = append(wl, wordList[i+1:]...)
		subans := findLadders(str, endWord, wl)
		if len(subans) == 0 {
			continue
		}
		if len(ans) != 0 && len(subans[0])+1 < len(ans[0]) {
			ans = ans[:0]
		}
		if len(ans) != 0 {
			subLen := len(subans[0]) + 1
			ansLen := len(ans[0])
			if subLen < ansLen {
				ans = ans[:0]
			} else if subLen > ansLen {
				subans = subans[:0]
			}
		}
		for _, sa := range subans {
			sa = append([]string{beginWord}, sa...)
			ans = append(ans, sa)
		}
	}
	return ans
}

// @lc code=end
