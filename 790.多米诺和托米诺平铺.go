/*
 * @lc app=leetcode.cn id=790 lang=golang
 *
 * [790] 多米诺和托米诺平铺
 */

// @lc code=start
func numTilings(N int) int {
	mod := 1_000_000_007
	dp := [4]int{1, 0, 0, 0}
	for i := 0; i < N; i++ {
		dp[0], dp[1], dp[2], dp[3] =
			dp[3]+dp[0], dp[0]+dp[2], dp[0]+dp[1], dp[0]+dp[1]+dp[2]
		dp[0] %= mod
		dp[1] %= mod
		dp[2] %= mod
		dp[3] %= mod
	}
	return dp[0]
}

// @lc code=end
