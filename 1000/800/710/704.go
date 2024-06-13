package leetcode_710

func search_recursion(nums []int, target int) int {
	return binarySearch(nums, 0, len(nums)-1, target)
}

func binarySearch(nums []int, start, end, target int) int {
	if start > end {
		return -1
	}

	mid := start + (end-start)/2
	midVal := nums[mid]

	if midVal == target {
		return mid
	} else if midVal < target {
		return binarySearch(nums, mid+1, end, target)
	} else {
		return binarySearch(nums, start, mid-1, target)
	}
}

func search_loop(nums []int, target int) int {
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
