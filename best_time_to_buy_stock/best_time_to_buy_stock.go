package main

import "fmt"

func maxProfit(prices []int) int {
	minPrice := prices[0]
	profit := 0
	for _, price := range prices {
		if price < minPrice {
			minPrice = price
		}
		profit = max(profit, price-minPrice)
	}
	return profit

}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}

func main() {
	a := []int{7, 1, 5, 3, 6, 4}
	fmt.Println(maxProfit(a) == 5)
	a = []int{7, 6, 4, 3, 1}
	fmt.Println(maxProfit(a) == 0)
}
