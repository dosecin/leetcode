/*
 * @lc app=leetcode.cn id=667 lang=golang
 *
 * [667] 优美的排列 II
 */

// @lc code=start
func constructArray(n int, k int) []int {
	if n == 0 || n <= k {
		return []int{}
	}
	ans := []int{}
	lef, rig := 0, n+1
	bAdd := true
	genFunc := func(bAdd bool) func() int {
		if bAdd {
			return func() int {
				lef++
				return lef
			}
		}
		return func() int {
			rig--
			return rig
		}
	}
	ansFunc := genFunc(bAdd)
	for i := 0; i < n; i++ {
		ans = append(ans, ansFunc())
		k--
		if k > 0 {
			bAdd = !bAdd
			ansFunc = genFunc(bAdd)
		}
	}
	return ans
}

// @lc code=end
