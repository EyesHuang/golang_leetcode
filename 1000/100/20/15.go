package leetcode_20

import "sort"

// Time complexity: O(n^2)
// Space complexity: O(n)
func threeSum_twoPointers(nums []int) [][]int {
	sort.Ints(nums)

	res := make([][]int, 0, len(nums))

	for i, val := range nums {
		// skip positive intergers
		if val > 0 {
			break
		}

		// skip duplicates
		if i > 0 && val == nums[i-1] {
			continue
		}

		twoSum_twoPointers(nums, i, &res)
	}

	return res
}

func twoSum_twoPointers(nums []int, i int, res *[][]int) {
	l, r := i+1, len(nums)-1

	for l < r {
		sum := nums[i] + nums[l] + nums[r]

		if sum < 0 {
			l++
		} else if sum > 0 {
			r--
		} else {
			*res = append(*res, []int{nums[i], nums[l], nums[r]})
			l++
			r--

			// skip duplicates
			for l < r && nums[l] == nums[l-1] {
				l++
			}
		}
	}
}

// Time complexity: O(n^2)
// Space complexity: O(n)
func threeSum_hashSet(nums []int) [][]int {
	sort.Ints(nums)
	res := make([][]int, 0, len(nums))

	for i, val := range nums {
		if val > 0 {
			break
		}

		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		twoSum_hashSet(nums, i, &res)
	}

	return res
}

func twoSum_hashSet(nums []int, i int, res *[][]int) {
	type void struct{}
	var member void

	m := make(map[int]void)

	for j := i + 1; j < len(nums); j++ {
		complement := -nums[i] - nums[j]
		if _, ok := m[complement]; ok {
			*res = append(*res, []int{nums[i], nums[j], complement})

			for j+1 < len(nums) && nums[j] == nums[j+1] {
				j++
			}
		}
		m[nums[j]] = member
	}
}
