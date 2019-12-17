/*
 * @lc app=leetcode.cn id=126 lang=golang
 *
 * [126] 单词接龙 II
 */

// @lc code=start
func findLadders(beginWord string, endWord string, wordList []string) [][]string {
	wordMap := map[string]int{}
	for _, word := range wordList {
		wordMap[word] = 1
	}

	if _, ok := wordMap[endWord]; !ok {
		return [][]string{}
	}

	getNext := func(beginWord string, exclude map[string]int) []string {
		ans := []string{}
		beginBytes := []byte(beginWord)
		for i := range beginBytes {
			oldByte := beginBytes[i]
			for b := byte('a'); b <= 'z'; b++ {
				if beginBytes[i] == b {
					continue
				}

				beginBytes[i] = b
				nextStr := string(beginBytes)
				if _, ok := exclude[nextStr]; ok {
					continue
				}
				if _, ok := wordMap[nextStr]; ok {
					ans = append(ans, nextStr)
				}
			}
			beginBytes[i] = oldByte
		}
		return ans
	}

	type solution struct {
		beginWord string
		list      []string
	}

	ans := [][]string{}
	inPath := map[string]int{beginWord: 1}
	solutions := []solution{
		solution{beginWord: beginWord, list: []string{beginWord}},
	}
	for len(ans) == 0 && len(solutions) > 0 {
		newSolutions := []solution{}
		currInPath := map[string]int{}
		for _, sol := range solutions {
			exclude := map[string]int{}
			for _, word := range sol.list {
				exclude[word] = 1
			}
			nexts := getNext(sol.beginWord, exclude)
			for _, next := range nexts {
				if _, ok := inPath[next]; ok {
					continue
				}
				if next == endWord {
					list := make([]string, len(sol.list), len(sol.list)+1)
					copy(list, sol.list)
					list = append(list, next)
					ans = append(ans, list)
					continue
				}
				currInPath[next] = 1
				newSol := solution{
					beginWord: next,
					list:      make([]string, len(sol.list), len(sol.list)+1),
				}
				copy(newSol.list, sol.list)
				newSol.list = append(newSol.list, next)
				newSolutions = append(newSolutions, newSol)
			}
		}
		solutions = newSolutions
		for k := range currInPath {
			inPath[k] = 1
		}
	}
	return ans
}

// @lc code=end
