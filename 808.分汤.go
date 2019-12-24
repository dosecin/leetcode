/*
 * @lc app=leetcode.cn id=808 lang=golang
 *
 * [808] 分汤
 */

// @lc code=start
func soupServings(N int) float64 {
	if N >= 5000 {
		return 1.0
	}
	if N%25 > 0 {
		N = N/25 + 1
	} else {
		N /= 25
	}
	M := func(a int) int {
		if a < 0 {
			return 0
		}
		return a
	}
	memo := make([][]float64, N+1)
	for i := 0; i <= N; i++ {
		memo[i] = make([]float64, N+1)
	}
	for i := 0; i <= N; i++ {
		ans := 0.0
		for j := 0; j <= N; j++ {
			if i == 0 && j == 0 {
				ans = 0.5
			} else if i == 0 {
				ans = 1.0
			} else if j != 0 {
				ans = 0.25 * (memo[M(i-4)][M(j)] + memo[M(i-3)][M(j-1)] + memo[M(i-2)][M(j-2)] + memo[M(i-1)][M(j-3)])
			}
			memo[i][j] = ans
		}
	}

	return memo[N][N]
}

// @lc code=end
