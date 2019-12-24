/*
 * @lc app=leetcode.cn id=798 lang=golang
 *
 * [798] 得分最高的最小轮调
 */

// @lc code=start
func bestRotation(A []int) int {
	scores := make([]int, len(A))
	for i := 0; i < len(A); i++ {
		lef, rig := (i-A[i]+1+len(A))%len(A), (i+1)%len(A)
		scores[lef]--
		scores[rig]++
	}
	maxScores, ans, score := -len(A), 0, 0
	for i, s := range scores {
		score += s
		if score > maxScores {
			maxScores = score
			ans = i
		}
	}
	return ans
}

// @lc code=end
