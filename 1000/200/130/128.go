package leetcode_130

import "sort"

// Time complexity: O(n)
// Space complexity: O(n)
func longestConsecutive_hashSet(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	numSet := make(map[int]bool)

	for _, num := range nums {
		numSet[num] = true
	}

	longest := 0

	for num := range numSet {
		if !numSet[num-1] {
			length := 1

			for numSet[num+length] {
				length++
			}

			if length > longest {
				longest = length
			}
		}
	}

	return longest
}

// Time complexity: O(n*log(n))
// Space complexity: O(log(n))
func longestConsecutive_sorting(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	sort.Ints(nums)
	length, longest := 1, 1

	for i := 1; i < len(nums); i++ {
		if nums[i] != nums[i-1] {
			if nums[i] == nums[i-1]+1 {
				length++
			} else {
				length = 1
			}

			if length > longest {
				longest = length
			}
		}

	}

	return longest
}
