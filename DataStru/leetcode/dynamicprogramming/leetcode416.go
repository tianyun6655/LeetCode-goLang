package dynamicprogramming

import "fmt"

func CanPartition(nums []int) bool {
	sum:=0
	for _,a:=range nums{
		sum+=a
	}
	if sum%2!=0{
		return false
	}
	length:=len(nums)

	sum = sum/2

	dp:=make([][]bool,length+1)

	for i:=0;i<length+1;i++{
		dp[i] =make([]bool,sum+1)
	}

	dp[0][0]=true
	for i:=1;i<length+1;i++{
		dp[i][0]=true
	}
	for j:=1;j<sum+1;j++{
		dp[0][j]=false
	}

	for i:=1;i<length+1;i++{
		for j:=1;j<sum+1;j++{
			dp[i][j]=dp[i-1][j]
			if j>=nums[i-1]{
				dp[i][j]=dp[i][j]||dp[i-1][j-nums[i-1]]
			}
		}
	}
	fmt.Println(dp)
	return dp[length][sum]
}
