package coin_change

import "math"

// CoinChangeBruteForce
// Time Complexity: O(S^n) (S is the amount, n is the number of denominations)
// Space Complexity: O(S)
func CoinChangeBruteForce(coins []int, amount int) int {
	if amount == 0 {
		return 0
	} else if amount < 0 {
		return -1
	}

	res := math.MaxInt
	for _, coin := range coins {
		subProblem := CoinChangeBruteForce(coins, amount-coin)
		if subProblem == -1 {
			continue
		}

		res = min(res, subProblem+1)
	}

	if res == math.MaxInt {
		return -1
	} else {
		return res
	}
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

// CoinChangeTopDown
// Time Complexity: O(S*n) (S is the amount, n is the number of denominations)
// Space Complexity: O(S) (S is the amount to change)
var memo []int

func CoinChangeTopDown(coins []int, amount int) int {
	memo = make([]int, amount+1)
	return recursionHelper(coins, amount)
}

func recursionHelper(coins []int, remain int) int {
	if remain < 0 {
		return -1
	} else if remain == 0 {
		return 0
	}

	if memo[remain] != 0 {
		return memo[remain]
	}

	minCount := math.MaxInt
	for _, coin := range coins {
		count := recursionHelper(coins, remain-coin)
		if count == -1 {
			continue
		}
		minCount = min(minCount, count+1)
	}

	if minCount == math.MaxInt {
		memo[remain] = -1
	} else {
		memo[remain] = minCount
	}

	return memo[remain]
}

// CoinChangeBottomUp
// Time Complexity: O(S*n) (S is the amount, n is the number of denominations)
// Space Complexity: O(S) (S is the amount to change)
func CoinChangeBottomUp(coins []int, amount int) int {
	dp := make([]int, amount+1)
	for i := 0; i < amount+1; i++ {
		dp[i] = amount + 1
	}
	dp[0] = 0
	for i := 1; i <= amount; i++ {
		for j := 0; j < len(coins); j++ {
			if coins[j] <= i {
				dp[i] = min(dp[i], dp[i-coins[j]]+1)
			}
		}
	}

	if dp[amount] == amount+1 {
		return -1
	} else {
		return dp[amount]
	}
}
