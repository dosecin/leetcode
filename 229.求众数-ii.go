/*
 * @lc app=leetcode.cn id=229 lang=golang
 *
 * [229] 求众数 II
 */

// @lc code=start
func majorityElement(nums []int) []int {
	pre := [2]int{}
	cnt := [2]int{}
	for _, num := range nums {
		if cnt[0] == 0 {
			pre[0] = num
			cnt[0] = 1
			continue
		}
		if num == pre[0] {
			cnt[0]++
			continue
		}
		if cnt[1] == 0 {
			pre[1] = num
			cnt[1] = 1
			continue
		}
		if num == pre[1] {
			cnt[1]++
			continue
		}
		cnt[0]--
		cnt[1]--
		if cnt[0] == 0 && cnt[1] > 0 {
			cnt[0] = cnt[1]
			pre[0] = pre[1]
			cnt[1] = 0
		}
	}
	ans := []int{}
	limit := len(nums) / 3
	cnt = [2]int{}
	for _, num := range nums {
		if num == pre[0] {
			cnt[0]++
		} else if num == pre[1] {
			cnt[1]++
		}
	}
	if cnt[0] > limit {
		ans = append(ans, pre[0])
	}
	if cnt[1] > limit {
		ans = append(ans, pre[1])
	}
	return ans
}

// @lc code=end
