/*
 * @lc app=leetcode.cn id=57 lang=golang
 *
 * [57] 插入区间
 */

// @lc code=start
func insert(intervals [][]int, newInterval []int) [][]int {
	index := []int{-1, -1}
	for i, nIntv := range newInterval {
		l, r := 0, len(intervals)-1
		for l <= r {
			m := l + (r-l)>>1
			if nIntv < intervals[m][0] {
				r = m - 1
			} else if nIntv > intervals[m][1] {
				l = m + 1
			} else {
				l = m
				break
			}
		}
		index[i] = l
	}
	l, r := index[0], index[1]
	if l == r {
		if l >= len(intervals) {
			intervals = append(intervals, newInterval)
		} else if newInterval[1] < intervals[l][0] {
			intervals = append(intervals, newInterval)
			copy(intervals[l+1:], intervals[l:])
			intervals[l] = newInterval
		} else {
			if newInterval[0] < intervals[l][0] {
				intervals[l][0] = newInterval[0]
			}
		}
	} else {
		if newInterval[0] < intervals[l][0] {
			intervals[l][0] = newInterval[0]
		}
		intervals[l][1] = newInterval[1]

		if r >= len(intervals) {
			intervals = intervals[:l+1]
		} else if newInterval[1] < intervals[r][0] {
			n := copy(intervals[l+1:], intervals[r:])
			intervals = intervals[:l+n+1]
		} else {
			intervals[l][1] = intervals[r][1]
			n := copy(intervals[l+1:], intervals[r+1:])
			intervals = intervals[:l+n+1]
		}
	}
	return intervals
}

// @lc code=end
