package climbing_stairs

func climbStairs_BruteForce(n int) int {
	return climb_stairs(0, n)
}

// i: current step
// n: destination step
func climb_stairs(i, n int) int {
	if i == n {
		return 1
	}

	if i > n {
		return 0
	}

	return climb_stairs(i+1, n) + climb_stairs(i+2, n)
}

func climbStairs_Memo(n int) int {
	memo := make([]int, n+1)
	return climb_stairs_memo(0, n, memo)
}

// i: current step
// n: destination step
func climb_stairs_memo(i, n int, memo []int) int {
	if i == n {
		return 1
	}

	if i > n {
		return 0
	}

	if memo[i] > 0 {
		return memo[i]
	}

	memo[i] = climb_stairs_memo(i+1, n, memo) + climb_stairs_memo(i+2, n, memo)

	return memo[i]
}

func climbStairs_DP(n int) int {
	if n == 1 {
		return 1
	}
	dp := make([]int, n+1)
	dp[1] = 1
	dp[2] = 2
	for i := 3; i <= n; i++ {
		dp[i] = dp[i-1] + dp[i-2]
	}
	return dp[n]
}
