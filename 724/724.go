package find_pivot_index

func PivotIndex(nums []int) int {
	length := len(nums)

	for i := 0; i < length; i++ {
		sumLeft, sumRight := 0, 0
		for j := i + 1; j < length; j++ {
			sumRight += nums[j]
		}

		if i > 0 {
			for k := 0; k < i; k++ {
				sumLeft += nums[k]
			}
		}

		if sumRight == 0 && i == 0 {
			return 0
		} else if sumLeft == sumRight {
			return i
		}
	}

	return -1
}

func PivotIndexS2(nums []int) int {
	sum, leftSum := 0, 0
	for _, num := range nums {
		sum += num
	}

	for i := 0; i < len(nums); i++ {
		if leftSum == sum-leftSum-nums[i] {
			return i
		}
		leftSum += nums[i]
	}

	return -1
}
