package leetcode

import "fmt"

/*
Follow up for "Unique Paths":

Now consider if some obstacles are added to the grids. How many unique paths would there be?

An obstacle and empty space is marked as 1 and 0 respectively in the grid.

For example,
There is one obstacle in the middle of a 3x3 grid as illustrated below.

[
  [0,0,0],
  [0,1,0],
  [0,0,0]
]
The total number of unique paths is 2.

Note: m and n will be at most 100.
*/

func UniquePathsWithObstacles(obstacleGrid [][]int) int {
	if len(obstacleGrid)==1||len(obstacleGrid[0])==0{
		if obstacleGrid[0][0]==1 {
			return 0
		}
		return 1
	}

	dp := make([][]int, len(obstacleGrid))

	// 创建二维Slice
	for i := range dp {
		subArray := make([]int, len(obstacleGrid[i]))
		dp[i] = subArray
	}

	for i:=0;i<len(obstacleGrid);i++{
		for j:=0;j<len(obstacleGrid[i]);j++{
			if obstacleGrid[i][j]== 1{
			 dp[i][j] = 0
			}else if i==0&&j==0{
				dp[i][j]=0
			} else if i==0 &&j!=0{
				dp[i][j]=1
			}else if j==0 &&i!=0{
				dp[i][j]=1
			}else {
				dp[i][j] = dp[i-1][j]+dp[i][j-1]

			}
		}
	}


	return dp[len(obstacleGrid)-1][len(obstacleGrid[0])-1]
}