package dynamicprogramming

import "fmt"

/*
You are given coins of different denominations and a total amount of money amount. Write a function to compute the fewest number of coins that you need to make up that amount. If that amount of money cannot be made up by any combination of the coins, return -1.

Example 1:
coins = [1, 2, 5], amount = 11
return 3 (11 = 5 + 5 + 1)

Example 2:
coins = [2], amount = 3
return -1.
*/

func CoinChange(coins []int, amount int) int {
	dp:=make([]int,amount+1)
	for i:=1;i<=amount;i++{
		min := 2147483647

		for j:=0;j<len(coins);j++{
			if i == coins[j]{
				min =1
				break
			}else{
				if i-coins[j]<0{
					continue
				}else if dp[coins[j]]==-1 ||dp[i-coins[j]]==-1{
					continue
				}
				temp :=dp[i-coins[j]]+dp[coins[j]]
				if temp < min{
					min=temp
				}
			}
		}

		if min==2147483647{
		dp[i] = -1
		}else{
			dp[i]=min
		}



	}
    fmt.Println(dp)
	return dp[amount]
}


