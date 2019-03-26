/*
 * @lc app=leetcode.cn id=53 lang=golang
 *
 * [53] 最大子序和
 *
 * https://leetcode-cn.com/problems/maximum-subarray/description/
 *
 * algorithms
 * Easy (43.50%)
 * Total Accepted:    44.1K
 * Total Submissions: 101.4K
 * Testcase Example:  '[-2,1,-3,4,-1,2,1,-5,4]'
 *
 * 给定一个整数数组 nums ，找到一个具有最大和的连续子数组（子数组最少包含一个元素），返回其最大和。
 *
 * 示例:
 *
 * 输入: [-2,1,-3,4,-1,2,1,-5,4],
 * 输出: 6
 * 解释: 连续子数组 [4,-1,2,1] 的和最大，为 6。
 *
 *
 * 进阶:
 *
 * 如果你已经实现复杂度为 O(n) 的解法，尝试使用更为精妙的分治法求解。
 *
 */

/*
 * 动态规划：sum(i)为以第i位数字为结尾的最大子序列和，那么以第i+1位数字为
 * 结尾的最大子序列和sum(i+1)=max(sum(i)+nums[i+1],nums[i+1])，因此算法
 * 仅需记录前一个sum(i)即可，同时记录一个整个数组的最大子序列和，每次对比即可
 */
func maxSubArray(nums []int) int {
	sum, max := nums[0], nums[0]
	for i := 1; i < len(nums); i++ {
		sum += nums[i]
		if sum < nums[i] {
			sum = nums[i]
		}
		if sum > max {
			max = sum
		}
	}
	return max
}
