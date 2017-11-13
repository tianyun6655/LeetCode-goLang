package dynamicprogramming

import (
	"math"
	"fmt"
)

func MaximalSquare(matrix [][]byte) int {

	length:=len(matrix)
	dp:=make([][]int,length)
	max:=0
	for i:=0;i<length;i++{
		dp[i]=make([]int,len(matrix[i]))

	}

	for i:=0;i<len(matrix[0]);i++{
		dp[0][i]=int(matrix[0][i])
		if max<dp[0][i]{
			max=dp[0][i]
		}
	}

	for j:=0;j<length;j++{
		dp[j][0]=int(matrix[j][0])
		if max<dp[j][0]{
			max=dp[j][0]
		}
	}

	for i:=1;i<length;i++{
		for j:=1;j<len(matrix[i]);j++{
			if matrix[i][j]!=0{
				dp[i][j]=int(math.Min(math.Min(float64(dp[i-1][j-1]),float64(dp[i][j-1])),float64(dp[i-1][j])))+1
			}else{
				dp[i][j]=1
			}

			if max<dp[i][j]{
				max=dp[i][j]
			}

		}
	}
	fmt.Println(dp)
	return max*max

}