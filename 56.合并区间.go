/*
 * @lc app=leetcode.cn id=56 lang=golang
 *
 * [56] 合并区间
 */

// @lc code=start
func merge(intervals [][]int) [][]int {
	ans := [][]int{}
	dich := func(base int) int {
		l, r := 0, len(ans)-1
		for l <= r {
			m := l + (r-l)>>1
			if base < ans[m][0] {
				r = m - 1
			} else if base > ans[m][1] {
				l = m + 1
			} else {
				l = m
				break
			}
		}
		return l
	}
	for _, intv := range intervals {
		l0 := dich(intv[0])
		l1 := dich(intv[1])
		if l0 != l1 {
			if intv[0] < ans[l0][0] {
				ans[l0][0] = intv[0]
			}
			ans[l0][1] = intv[1]
			if l1 >= len(ans) {
				ans = ans[:l0+1]
			} else if intv[1] < ans[l1][0] {
				n := copy(ans[l0+1:], ans[l1:])
				ans = ans[:l0+n+1]
			} else {
				ans[l0][1] = ans[l1][1]
				n := copy(ans[l0+1:], ans[l1+1:])
				ans = ans[:l0+1+n]
			}
		} else {
			if l0 >= len(ans) {
				ans = append(ans, []int{intv[0], intv[1]})
			} else if intv[1] < ans[l0][0] {
				tail := []int{intv[0], intv[1]}
				ans = append(ans, tail)
				copy(ans[l0+1:], ans[l0:])
				ans[l0] = tail
			} else {
				if intv[0] < ans[l0][0] {
					ans[l0][0] = intv[0]
				}
			}
		}
	}
	return ans
}

// @lc code=end
