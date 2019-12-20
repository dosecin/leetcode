/*
 * @lc app=leetcode.cn id=714 lang=golang
 *
 * [714] 买卖股票的最佳时机含手续费
 */

// @lc code=start
func maxProfit(prices []int, fee int) int {
	if len(prices) == 0 {
		return 0
	}
	p0, p1 := 0, -prices[0]
	for _, p := range prices[1:] {
		p0, p1 = maxInt(p0, p1+p-fee), maxInt(p1, p0-p)
	}
	return p0
}

func maxInt(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// @lc code=end
