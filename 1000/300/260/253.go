package leetcode_260

import "sort"

// Time complexity: O(n*log(n))
// Space complexity: O(n)
func minMeetingRooms(intervals [][]int) int {
	n := len(intervals)
	startArr, endArr := make([]int, n), make([]int, n)

	for i := 0; i < n; i++ {
		startArr[i] = intervals[i][0]
		endArr[i] = intervals[i][1]
	}

	sort.Ints(startArr)
	sort.Ints(endArr)

	s, e, res, count := 0, 0, 0, 0

	for s < n {
		if startArr[s] < endArr[e] {
			count++
			s++
		} else {
			count--
			e++
		}

		if count > res {
			res = count
		}
	}

	return res
}
