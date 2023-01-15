package best_time_to_buy_and_sell_stock

import (
	"math"
)

//func main() {
//	var prices []int
//	r := maxProfit(prices)
//	fmt.Println(r)
//}

func maxProfitS1(prices []int) int {
	profit := 0
	minPrice := math.MaxInt

	for i := 0; i < len(prices); i++ {
		if prices[i] < minPrice {
			minPrice = prices[i]
		} else if prices[i]-minPrice > profit {
			profit = prices[i] - minPrice
		}
	}

	return profit
}
