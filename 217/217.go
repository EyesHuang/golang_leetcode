package contains_duplicate

// Time complexity: O(n^2)
// Space complexity: O(1)
func containsDuplicate_bruteForce(nums []int) bool {
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i] == nums[j] {
				return true
			}
		}
	}

	return false
}

// Time complexity: O(n)
// Space complexity: O(n)
func containsDuplicate_set(nums []int) bool {
	type void struct{}
	var member void

	set := make(map[int]void)

	for _, num := range nums {
		if _, ok := set[num]; ok {
			return true
		} else {
			set[num] = member
		}
	}

	return false
}
