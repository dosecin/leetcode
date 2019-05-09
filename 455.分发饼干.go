/*
 * @lc app=leetcode.cn id=455 lang=golang
 *
 * [455] 分发饼干
 */
func findContentChildren(g []int, s []int) int {
	sort.Ints(g)
	sort.Ints(s)
	res := 0
	i, j := 0, 0
    for i < len(g) && j < len(s) {
		if s[j] >= g[i] {
			res++
			i++
		}
		j++
	}
	return res
}

