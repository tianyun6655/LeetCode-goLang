package dynamicprogramming


func maxProfit(prices []int) int {
	minPrice:=prices[0]
	if len(prices)==1{
		return 0
	}
	maxProfit :=0
	if prices[1]-prices[0]>0{
		maxProfit=prices[1]-prices[0]
	}
	for i:=2;i<len(prices);i++{
		if  minPrice<prices[i-1]{
			minPrice = prices[i-1]
		}
		if prices[i]>maxProfit{
			currentPrices:=prices[i]-minPrice
			if maxProfit < currentPrices {
				maxProfit=currentPrices
			}
		}
	}
	return maxProfit
}