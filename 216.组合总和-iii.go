/*
 * @lc app=leetcode.cn id=216 lang=golang
 *
 * [216] 组合总和 III
 */

// @lc code=start
func combinationSum3(k int, n int) [][]int {
	ans := [][]int{}
	sol := []int{1}
	sum := 1
	prune := func() {
		for {
			sum -= sol[len(sol)-1]
			sol = sol[:len(sol)-1]
			if len(sol) == 0 {
				break
			}
			lastIdx := len(sol) - 1
			lastData := sol[lastIdx]
			if lastData < 9 {
				sol[lastIdx]++
				sum++
				break
			}
		}
	}
	for len(sol) > 0 {
		lastIdx := len(sol) - 1
		lastData := sol[lastIdx]
		if len(sol) == k {
			if sum == n {
				ans = append(ans, append([]int{}, sol...))
			}
			if lastData == 9 || sum >= n {
				prune()
			} else {
				sol[lastIdx]++
				sum++
			}
		} else {
			if lastData == 9 || sum >= n {
				prune()
			} else {
				sol = append(sol, lastData+1)
				sum += lastData + 1
			}
		}
	}
	return ans
}

// @lc code=end
