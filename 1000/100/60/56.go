package leetcode_60

import "sort"

// Time complexity: O(n*log(n))
// Space complexity: O(n)
func merge(intervals [][]int) [][]int {
	if len(intervals) == 0 || len(intervals) == 1 {
		return intervals
	}

	sort.Slice(intervals, func(i, j int) bool { return intervals[i][0] < intervals[j][0] })

	res := make([][]int, 0, len(intervals))
	res = append(res, intervals[0])

	for i := 1; i < len(intervals); i++ {
		currStart, currEnd := intervals[i][0], intervals[i][1]
		lastEnd := res[len(res)-1][1]

		if lastEnd >= currStart {
			res[len(res)-1][1] = max(lastEnd, currEnd)
		} else {
			res = append(res, intervals[i])
		}
	}

	return res

}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
