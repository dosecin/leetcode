import (
	"math"
	"sort"
)

/*
 * @lc app=leetcode.cn id=857 lang=golang
 *
 * [857] 雇佣 K 名工人的最低成本
 */

// @lc code=start
func mincostToHireWorkers(quality []int, wage []int, K int) float64 {
	type cost struct {
		value float64
		qua   float64
	}
	costs := make([]cost, len(quality))
	for i, qua := range quality {
		costs[i] = cost{
			value: float64(wage[i]) / float64(qua),
			qua:   float64(qua),
		}
	}
	sort.Slice(costs, func(i, j int) bool {
		return costs[i].value <= costs[j].value
	})
	selecteds := []cost{}
	up := func() {
		i := len(selecteds) - 1
		for {
			p := (i - 1) / 2
			if p < 0 {
				break
			}
			if selecteds[i].qua <= selecteds[p].qua {
				break
			}
			selecteds[i], selecteds[p] = selecteds[p], selecteds[i]
			i = p
		}
	}
	down := func() {
		i := 0
		for {
			lef := 2*i + 1
			if lef >= len(selecteds) {
				break
			}
			maxIdx := lef
			if rig := lef + 1; rig < len(selecteds) && selecteds[rig].qua > selecteds[lef].qua {
				maxIdx = rig
			}
			if selecteds[i].qua >= selecteds[maxIdx].qua {
				break
			}
			selecteds[i], selecteds[maxIdx] = selecteds[maxIdx], selecteds[i]
			i = maxIdx
		}
	}
	sumQuality := 0.0
	ans := math.MaxFloat64
	for _, c := range costs {
		if len(selecteds) < K {
			sumQuality += c.qua
			selecteds = append(selecteds, c)
			up()
			if len(selecteds) == K {
				ans = sumQuality * c.value
			}
			continue
		}
		sumQuality += c.qua - selecteds[0].qua
		selecteds[0] = c
		down()
		currAns := sumQuality * c.value
		if ans > currAns {
			ans = currAns
		}
	}
	return ans
}

// @lc code=end
