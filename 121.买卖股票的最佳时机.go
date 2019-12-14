/*
 * @lc app=leetcode.cn id=121 lang=golang
 *
 * [121] 买卖股票的最佳时机
 */

// @lc code=start
import "math"

func maxProfit(prices []int) int {
	maxInt := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	// f0为当前不持股状态，f1为当前持股状态
	f0, f1 := 0, math.MinInt64
	for _, p := range prices {
		f0 = maxInt(f0, f1+p)
		f1 = maxInt(f1, -p)
	}
	return f0
}

// @lc code=end
