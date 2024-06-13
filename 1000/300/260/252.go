package leetcode_260

import "sort"

// Time complexity: O(n^2)
// Space complexity: O(1)
func canAttendMeetings_bruteForce(intervals [][]int) bool {
	for i := 0; i < len(intervals); i++ {
		for j := i + 1; j < len(intervals); j++ {
			if overlap(intervals[i], intervals[j]) {
				return false
			}
		}
	}

	return true
}

func overlap(interval1, interval2 []int) bool {
	return interval1[0] < interval2[1] && interval1[1] > interval2[0]
}

// Time complexity: O(nlogn)
// Space complexity: O(1)
func canAttendMeetings_Sorting(intervals [][]int) bool {
	sort.Slice(intervals, func(i, j int) bool { return intervals[i][0] < intervals[j][0] })

	for i := 0; i < len(intervals)-1; i++ {
		if intervals[i][1] > intervals[i+1][0] {
			return false
		}
	}

	return true
}
