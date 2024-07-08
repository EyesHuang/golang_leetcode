package leetcode_20

import "leetcode/util"

// Time complexity: O(n^2)
// Space complexity: O(1)
func maxArea_bruteForce(height []int) int {
	res := 0
	n := len(height)

	for left := 0; left < n; left++ {
		for right := left + 1; right < n; right++ {
			area := util.Min(height[left], height[right]) * (right - left)
			res = util.Max(res, area)
		}
	}
	return res
}

// Time complexity: O(n)
// Space complexity: O(1)
func maxArea_twoPointers(height []int) int {
	res, left, right := 0, 0, len(height)-1

	for left < right {
		area := util.Min(height[left], height[right]) * (right - left)
		res = util.Max(res, area)

		if height[left] < height[right] {
			left++
		} else {
			right--
		}
	}
	return res
}
