package togroup

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
