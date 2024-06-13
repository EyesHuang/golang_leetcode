package leetcode_10

// Time complexity: O(n^2)
// Space complexity: O(1)
func twoSum_bruteForce(nums []int, target int) []int {
	sum := 0

	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			sum = nums[i] + nums[j]

			if sum == target {
				return []int{i, j}
			}
		}
	}

	return nil
}

// Time complexity: O(n)
// Space complexity: O(n)
func twoSum_hashMap(nums []int, target int) []int {
	m := make(map[int]int)

	for i := 0; i < len(nums); i++ {
		current := nums[i]

		if val, ok := m[target-current]; ok {
			return []int{val, i}
		} else {
			m[current] = i
		}
	}

	return nil
}
