func maxProfit(prices []int) int {
    maxProfit := 0
    buyPrice := math.MaxInt

    for _, price := range prices {
        buyPrice = min(buyPrice, price)
        profit := price - buyPrice
        maxProfit = max(maxProfit, profit)
    }

    return maxProfit
}
