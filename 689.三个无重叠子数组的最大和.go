/*
 * @lc app=leetcode.cn id=689 lang=golang
 *
 * [689] 三个无重叠子数组的最大和
 */

// @lc code=start
func maxSumOfThreeSubarrays(nums []int, k int) []int {
	/*
		var m1s0, m1s1, m2s0, m2s1, m3s0, m3s1 int
		ans1, ans2, ans3 := []int{0}, []int{}, []int{}
		m2s0k, m1s0k := make([]int, k), make([]int, k)
		ans1k, ans2k := make([][1]int, k), make([][2]int, k)
		sumk := 0
		for i := 0; i < len(nums); i++ {
			beginIdx := 0
			sumk += nums[i]
			if i >= k {
				beginIdx = i - k + 1
				sumk -= nums[i-k]
				m1s0k[i%k] = m1s0
				copy(ans1k[i%k][:], ans1)
				m2s0k[i%k] = m2s0
				copy(ans2k[i%k][:], ans2)
			}
			kIdx := (i + 1) % k
			m3s1 = sumk + m2s0k[kIdx]
			if m3s1 > m3s0 {
				ans3 = append(ans2k[kIdx][:], beginIdx)
				m3s0 = m3s1
			}

			m2s1 = sumk + m1s0k[kIdx]
			if m2s1 > m2s0 {
				ans2 = append(ans1k[kIdx][:], beginIdx)
				m2s0 = m2s1
			}

			m1s1 = sumk
			if m1s1 > m1s0 {
				ans1[0] = beginIdx
				m1s0 = m1s1
			}
		}
		return ans3
	*/
	// 官方解法
	sums := []int{}
	{
		sum := 0
		for i := range nums {
			sum += nums[i]
			if i >= k {
				sum -= nums[i-k]
			}
			sums = append(sums, sum)
		}
		sums = sums[k-1:]
	}
	lef, rig := make([]int, len(sums)), make([]int, len(sums))
	maxLef, maxRig := 0, len(sums)-1
	for i := range sums {
		idxLef, idxRig := i, len(sums)-1-i
		if sums[maxLef] < sums[idxLef] {
			lef[idxLef] = idxLef
			maxLef = idxLef
		} else {
			lef[idxLef] = maxLef
		}
		if sums[maxRig] <= sums[idxRig] {
			rig[idxRig] = idxRig
			maxRig = idxRig
		} else {
			rig[idxRig] = maxRig
		}
	}
	ans := make([]int, 3)
	max := 0
	for i := range sums {
		if i+2*k >= len(sums) {
			break
		}
		l, m, r := lef[i], i+k, rig[i+2*k]
		sum := sums[l] + sums[m] + sums[r]
		if sum > max {
			max = sum
			ans[0], ans[1], ans[2] = l, m, r
		}
	}
	return ans
}

// @lc code=end
