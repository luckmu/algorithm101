package greedy

import (
	"sort"
)

// 435
// 任意2个pair，只要有重叠，必定移除一个
// 若是乱序，无法确定移除哪一个
// e.g.
// [1,2], [1,3]
// [1,4], [2,5]
// 按 interval[0] 排序，保留 interval[1]小的（删除 interval[1]大的）
func isoverlap(intv1, intv2 []int) bool {
	if intv2[0] >= intv1[1] || intv1[0] >= intv2[1] {
		return false
	}
	return true
}

func eraseOverlapIntervals(intervals [][]int) int {
	if len(intervals) <= 1 {
		return 0
	}
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	var count int
	for i, j := 0, 1; i < len(intervals) && j < len(intervals); {
		if isoverlap(intervals[i], intervals[j]) {
			count++
			if intervals[j][1] > intervals[i][1] {
				j++
			} else {
				i, j = j, j+1
			}
		} else {
			i, j = j, j+1
		}
	}
	return count
}
