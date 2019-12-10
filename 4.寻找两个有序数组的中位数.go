/*
 * @lc app=leetcode.cn id=4 lang=golang
 *
 * [4] 寻找两个有序数组的中位数
 */

// @lc code=start
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	if len(nums1) > len(nums2) {
		return findMedianSortedArrays(nums2, nums1)
	}
	n, m := len(nums1), len(nums2)
	iMin, iMax := 0, n
	for iMin <= iMax {
		i := (iMin + iMax) >> 1
		j := (n+m+1)>>1 - i
		if j != 0 && i != n && nums2[j-1] > nums1[i] {
			iMin = i + 1
		} else if i != 0 && j != m && nums1[i-1] > nums2[j] {
			iMax = i - 1
		} else {
			maxLeft := 0
			if i == 0 {
				maxLeft = nums2[j-1]
			} else if j == 0 {
				maxLeft = nums1[i-1]
			} else {
				maxLeft = max(nums1[i-1], nums2[j-1])
			}
			if (n+m)%2 != 0 {
				return float64(maxLeft)
			}

			minRight := 0
			if i == n {
				minRight = nums2[j]
			} else if j == m {
				minRight = nums1[i]
			} else {
				minRight = min(nums1[i], nums2[j])
			}
			return float64(maxLeft+minRight) / 2
		}
	}
	return 0.0
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// @lc code=end
