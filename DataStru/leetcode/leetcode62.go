package leetcode

import "fmt"

/*
A robot is located at the top-left corner of a m x n grid (marked 'Start' in the diagram below).

The robot can only move either down or right at any point in time. The robot is trying to reach the bottom-right corner of the grid (marked 'Finish' in the diagram below).

How many possible unique paths are there?
*/


func UniquePaths(m int, n int) int {

	dp := make([][]int, m)

	// 创建二维Slice
	for i := range dp {
		subArray := make([]int, n)
		dp[i] = subArray
	}
	for i:=0;i<m;i++{
		for j:=0;j<n;j++{
            if i==0&&j==0{
				dp[i][j]=0
			}else if i==0&&j!=0{
				dp[i][j]=1
			}else if j==0&&i!=0{
				dp[i][j] = 1
			}else{
				dp[i][j] = dp[i][j-1]+dp[i-1][j]
			}
		}
	}
	fmt.Println(dp)
	return dp[m-1][n-1]
}