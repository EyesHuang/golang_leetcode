package best_time_to_buy_and_sell_stock

import (
	"math"
)

//func main() {
//	prices := []int{7, 6, 4, 3, 1}
//	r := maxProfit(prices)
//	fmt.Println(r)
//}

func maxProfitS2(prices []int) int {
	gain := make([]int, len(prices)-1)

	for i := 1; i < len(prices); i++ {
		gain[i-1] = prices[i] - prices[i-1]
	}

	profit := maxSubArray(gain)
	if profit < 0 {
		profit = 0
	}
	return profit
}

func maxSubArray(nums []int) int {
	cMax := 0
	max := math.MinInt

	for i := 0; i < len(nums); i++ {
		cMax = cMax + nums[i]
		if cMax > max {
			max = cMax
		}

		if cMax < 0 {
			cMax = 0
		}
	}

	return max
}
