package binary_search

func Search(nums []int, target int) int {
	start, end, mid := 0, len(nums)-1, 0

	for start <= end {
		mid = start + (end-start)/2

		if nums[mid] == target {
			return mid
		} else if nums[mid] > target {
			end = mid - 1
		} else {
			start = mid + 1
		}
	}

	return -1
}
