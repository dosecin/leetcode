/*
 * @lc app=leetcode.cn id=718 lang=golang
 *
 * [718] 最长重复子数组
 */

// @lc code=start
func findLength(A []int, B []int) int {
	ans := 0
	for i := range A {
		tmpA := A[i:]
		cnt := 0
		j := 0
		for j < len(tmpA) && j < len(B) {
			if tmpA[j] == B[j] {
				cnt++
				if cnt > ans {
					ans = cnt
				}
			} else {
				cnt = 0
			}
			j++
		}
	}
	for i := range B {
		tmpB := B[i:]
		cnt := 0
		j := 0
		for j < len(tmpB) && j < len(A) {
			if tmpB[j] == A[j] {
				cnt++
				if cnt > ans {
					ans = cnt
				}
			} else {
				cnt = 0
			}
			j++
		}
	}
	return ans
}

// @lc code=end
