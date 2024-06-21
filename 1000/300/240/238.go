package leetcode_240

// Time complexity: O(n)
// Space complexity: O(n)
func productExceptSelf(nums []int) []int {
	length := len(nums)
	left := make([]int, length)
	right := make([]int, length)
	res := make([]int, length)

	left[0] = 1
	for i := 1; i < length; i++ {
		left[i] = nums[i-1] * left[i-1]
	}

	right[length-1] = 1
	for i := length - 2; i >= 0; i-- {
		right[i] = nums[i+1] * right[i+1]
	}

	for i := 0; i < length; i++ {
		res[i] = left[i] * right[i]
	}

	return res
}

// Time complexity: O(n)
// Space complexity: O(n)
func productExceptSelf_optimization(nums []int) []int {
	length := len(nums)
	res := make([]int, length)

	res[0] = 1
	for i := 1; i < length; i++ {
		res[i] = nums[i-1] * res[i-1]
	}

	right := 1
	for i := length - 1; i >= 0; i-- {
		res[i] *= right
		right *= nums[i]
	}

	return res
}
