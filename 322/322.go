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
