package leetcode_60

// Time complexity: O(n)
// Space complexity: O(n)
func insert_linear(intervals [][]int, newInterval []int) [][]int {
	res := [][]int{}
	i := 0

	// Add all intervals before the new interval
	for i < len(intervals) && intervals[i][1] < newInterval[0] {
		res = append(res, intervals[i])
		i++
	}

	// Merge all overlapping intervals with the new interval
	for i < len(intervals) && intervals[i][0] <= newInterval[1] {
		newInterval[0] = min(newInterval[0], intervals[i][0])
		newInterval[1] = max(newInterval[1], intervals[i][1])
		i++
	}
	res = append(res, newInterval)

	// Add all remaining intervals after the new interval
	for i < len(intervals) {
		res = append(res, intervals[i])
		i++
	}

	return res
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func insert_binarySearch(intervals [][]int, newInterval []int) [][]int {
	// If the intervals vector is empty, return a vector containing the newInterval
	if len(intervals) == 0 {
		return [][]int{newInterval}
	}

	n := len(intervals)
	target := newInterval[0]
	left, right := 0, n-1

	// Binary search to find the position to insert newInterval
	for left <= right {
		mid := (left + right) / 2
		if intervals[mid][0] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	// Insert newInterval at the found position
	result := make([][]int, 0)
	result = append(result, intervals[:left]...)
	result = append(result, newInterval)
	result = append(result, intervals[left:]...)

	// Merge overlapping intervals
	res := make([][]int, 0)
	for _, interval := range result {
		// If res is empty or there is no overlap, add the interval to the result
		if len(res) == 0 || res[len(res)-1][1] < interval[0] {
			res = append(res, interval)
			// If there is an overlap, merge the intervals by updating the end of the last interval in res
		} else {
			if res[len(res)-1][1] < interval[1] {
				res[len(res)-1][1] = interval[1]
			}
		}
	}

	return res
}
