/*
 * @lc app=leetcode.cn id=169 lang=golang
 *
 * [169] 求众数
 *
 * https://leetcode-cn.com/problems/majority-element/description/
 *
 * algorithms
 * Easy (58.32%)
 * Total Accepted:    31.4K
 * Total Submissions: 53.8K
 * Testcase Example:  '[3,2,3]'
 *
 * 给定一个大小为 n 的数组，找到其中的众数。众数是指在数组中出现次数大于 ⌊ n/2 ⌋ 的元素。
 *
 * 你可以假设数组是非空的，并且给定的数组总是存在众数。
 *
 * 示例 1:
 *
 * 输入: [3,2,3]
 * 输出: 3
 *
 * 示例 2:
 *
 * 输入: [2,2,1,1,1,2,2]
 * 输出: 2
 *
 *
 */
func majorityElement(nums []int) int {
	count, maj := 0, 0
	for _, v := range nums {
		if count == 0 {
			maj = v
			count = 1
		} else {
			if v == maj {
				count++
			} else {
				count--
			}
		}
	}
	return maj
}
