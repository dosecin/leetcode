/*
 * @lc app=leetcode.cn id=532 lang=golang
 *
 * [532] 数组中的K-diff数对
 */
func findPairs(nums []int, k int) int {
	if k < 0 {
		return 0
	}
	nm := map[int]int{}
	for _, n := range nums {
		nm[n]++
	}
	res := 0
	if k == 0 {
		for _, v := range nm {
			if v > 1 {
				res++
			}
		}
	} else {
		for v := range nm {
			x := v + k
			if _, ok := nm[x]; ok {
				res++
			}
		}
	}
	return res
}
