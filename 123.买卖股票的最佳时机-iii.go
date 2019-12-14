/*
 * @lc app=leetcode.cn id=123 lang=golang
 *
 * [123] 买卖股票的最佳时机 III
 */

// @lc code=start
import "math"

func maxProfit(prices []int) int {
	/*穷举
	once := func(prices []int) int {
		mp := 0
		currProfit := 0
		for i := 1; i < len(prices); i++ {
			currProfit += prices[i] - prices[i-1]
			if currProfit < 0 {
				currProfit = 0
			}
			if currProfit > mp {
				mp = currProfit
			}
		}
		return mp
	}
	mp := once(prices)
	end := len(prices) - 1
	for i := 2; i < end; i++ {
		cp := once(prices[:i]) + once(prices[i:])
		if cp > mp {
			mp = cp
		}
	}
	return mp
	*/
	maxInt := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	f2k0, f2k1 := 0, math.MinInt64
	f1k0, f1k1 := 0, math.MinInt64
	for _, p := range prices {
		f2k0 = maxInt(f2k0, f2k1+p)
		f2k1 = maxInt(f2k1, f1k0-p)
		f1k0 = maxInt(f1k0, f1k1+p)
		f1k1 = maxInt(f1k1, -p)
	}
	return f2k0
}

// @lc code=end
