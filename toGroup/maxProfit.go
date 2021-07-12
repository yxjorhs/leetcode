package togroup

/*maxProfit 最佳卖卖时机，只能买入，卖出1次 */
func maxProfit(prices []int) int {
	maxPro := 0
	minPri := -1

	for i := 0; i < len(prices); i++ {
		if minPri == -1 || minPri > prices[i] {
			minPri = prices[i]
			continue
		}

		if prices[i]-minPri > maxPro {
			maxPro = prices[i] - minPri
		}
	}

	return maxPro
}

/*maxProfit 最佳卖卖时机，可买入，卖出多次 */
func maxProfit2(prices []int) int {
	maxPro := 0
	cost := -1 // 买入成本

	for i := 0; i < len(prices); i++ {
		if cost == -1 || cost > prices[i] {
			cost = prices[i]
			continue
		}

		maxPro += prices[i] - cost
		cost = prices[i]
	}

	return maxPro
}
