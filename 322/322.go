package coin_change

import "math"

// CoinChangeBruteForce
// Time Complexity: O(S^n) (S is the amount, n is the number of coins)
// Space Complexity: O(n)
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
// Time Complexity: O(S*n) (S is the amount, n is the number of coins)
// Space Complexity: O(S) (S is the amount to change)
func CoinChangeTopDown(coins []int, amount int) int {
	if amount < 1 {
		return 0
	}
	return coinChange(coins, amount, make([]int, amount))
}

func coinChange(coins []int, rem int, count []int) int {
	if rem < 0 {
		return -1
	} else if rem == 0 {
		return 0
	}

	if count[rem-1] != 0 {
		return count[rem-1]
	}

	minVal := math.MaxInt
	for _, coin := range coins {
		res := coinChange(coins, rem-coin, count)
		if res >= 0 && res < minVal {
			minVal = 1 + res
		}
	}

	if minVal == math.MaxInt {
		return -1
	} else {
		return minVal
	}
}

// CoinChangeBottomUp
// Time Complexity: O(S*n) (S is the amount, n is the number of coins)
// Space Complexity: O(S) (S is the amount to change)
func CoinChangeBottomUp(coins []int, amount int) int {
	max := amount + 1
	dp := make([]int, max)
	for i := 0; i < max; i++ {
		dp[i] = max
	}
	dp[0] = 0
	for i := 1; i <= amount; i++ {
		for j := 0; j < len(coins); j++ {
			if coins[j] <= i {
				dp[i] = min(dp[i], dp[i-coins[j]]+1)
			}
		}
	}

	if dp[amount] > amount {
		return -1
	} else {
		return dp[amount]
	}
}
