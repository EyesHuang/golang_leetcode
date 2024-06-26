package leetcode_1480

func runningSum(nums []int) []int {
	length := len(nums)
	res := make([]int, length)
	res[0] = nums[0]

	for i := 1; i < length; i++ {
		res[i] = res[i-1] + nums[i]
	}

	return res
}
