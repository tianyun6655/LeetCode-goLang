package dynamicprogramming

import "fmt"

/*
There is an m by n grid with a ball. Given the start coordinate (i,j) of the ball, you can move the ball to adjacent cell or cross the grid boundary in four directions (up, down, left, right). However, you can at most move N times. Find out the number of paths to move the ball out of grid boundary. The answer may be very large, return it after mod 109 + 7.

Example 1:
Input:m = 2, n = 2, N = 2, i = 0, j = 0
Output: 6
Explanation:

*/


func FindPaths(m int, n int, N int, i int, j int) int {
   limit:=1000000007
	var dp [10][10][10] int
	for k:=1;k<=N;k++ {
		for i := 0; i < m; i++ {
			for j:=0;j<n;j++{
				if i==0{
					dp[k][i][j]+=1
				}else{
					dp[k][i][j]+=dp[k-1][i-1][j]
				}

				if i==m-1{
					dp[k][i][j]+=1
				}else{
					dp[k][i][j]+=dp[k-1][i+1][j]
				}

				if j==0{
					dp[k][i][j]+=1
				}else{
					dp[k][i][j]+=dp[k-1][i][j-1]
				}

				if j==n-1{
					dp[k][i][j]+=1
				}else{
					dp[k][i][j]+=dp[k-1][i][j+1]
				}

				dp[k][i][j]%=limit
			}
		}
	}
	fmt.Println(dp)
	return dp[N][i][j]
}