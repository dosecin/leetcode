/*
 * @lc app=leetcode.cn id=33 lang=golang
 *
 * [33] 搜索旋转排序数组
 */
func search(nums []int, target int) int {
	low, hig := 0, len(nums)-1
	for low <= hig {
		mid := low + (hig-low)>>1
		if nums[mid] == target {
			return mid
		}
		if nums[low] <= nums[mid] {
			if target < nums[mid] && target >= nums[low] {
				res := binarySearch(nums[low:mid], target)
				if res == -1 {
					return -1
				}
				return low + res
			}
			low = mid + 1
		} else {
			if target > nums[mid] && target <= nums[hig] {
				res := binarySearch(nums[mid+1:hig+1], target)
				if res == -1 {
					return -1
				}
				return mid + 1 + res
			}
			hig = mid - 1
		}
	}
	return -1
}

func binarySearch(nums []int, target int) int {
	low, hig := 0, len(nums)-1
	for low <= hig {
		mid := low + (hig-low)>>1
		switch {
		case nums[mid] == target:
			return mid
		case nums[mid] < target:
			low = mid + 1
		case nums[mid] > target:
			hig = mid - 1
		}
	}
	return -1
}
