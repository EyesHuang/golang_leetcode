package leetcode_440

import (
	"sort"

	"leetcode/util"
)

// Time complexity: O(n*log(n))
// Space complexity: O(n)
func eraseOverlapIntervals(intervals [][]int) int {
	sort.Slice(intervals, func(i, j int) bool { return intervals[i][0] < intervals[j][0] })

	prevEnd := intervals[0][1]
	res := 0

	for i := 1; i < len(intervals); i++ {
		start, end := intervals[i][0], intervals[i][1]

		if start >= prevEnd {
			prevEnd = end
		} else {
			res++
			prevEnd = util.Min(end, prevEnd)
		}
	}

	return res
}
