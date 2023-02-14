package fibonacci_number

// FibBruteforce
// Time Complexity: O(2^n)
func FibBruteforce(n int) int {
	if n == 1 || n == 2 {
		return 1
	}
	return FibBruteforce(n-1) + FibBruteforce(n-2)
}

// FibTopDown
// Time Complexity: O(n)
func FibTopDown(n int) int {
	memo := make([]int, n+1)
	return helper(memo, n)
}

func helper(memo []int, n int) int {
	if n == 0 || n == 1 {
		return n
	}

	if memo[n] != 0 {
		return memo[n]
	}

	memo[n] = helper(memo, n-1) + helper(memo, n-2)
	return memo[n]
}
