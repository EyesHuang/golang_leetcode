package meeting_rooms

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
