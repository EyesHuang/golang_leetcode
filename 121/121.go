package best_time_to_buy_and_sell_stock

import "math"

// Time complexity: O(n^2)
// Space complexity: O(1)
func maxProfit_bruteForce(prices []int) int {
	max := 0

	for i := 0; i < len(prices); i++ {
		for j := i + 1; j < len(prices); j++ {
			tmp := prices[j] - prices[i]

			if tmp > max {
				max = tmp
			}
		}
	}

	return max
}

// Time complexity: O(n)
// Space complexity: O(1)
func maxProfit_slidingWindow(prices []int) int {
	minPrice := math.MaxInt
	profit := 0

	for i := 0; i < len(prices); i++ {
		if prices[i] < minPrice {
			minPrice = prices[i]
		} else if currentProfit := prices[i] - minPrice; currentProfit > profit {
			profit = prices[i] - minPrice
		}
	}

	return profit
}
