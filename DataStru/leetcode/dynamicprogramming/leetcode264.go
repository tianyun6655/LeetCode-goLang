package dynamicprogramming

import "math"

/*
Write a program to find the n-th ugly number.

Ugly numbers are positive numbers whose prime factors only include 2, 3, 5. For example, 1, 2, 3, 4, 5, 6, 8, 9, 10, 12 is the sequence of the first 10 ugly numbers.

Note that 1 is typically treated as an ugly number, and n does not exceed 1690.

Credits:
Special thanks to @jianchao.li.fighter for adding this problem and creating all test cases.


*/

func NthUglyNumber(n int) int {
	if n==0{
		return 0
	}
	dp:=make([]float64,n)
	index2,index3,index5:=0,0,0
	factor2,factor3,factor5:=2.0,3.0,5.0

	for i:=1;i<len(dp);i++{
		min:=math.Min(math.Min(factor2,factor3),factor5)
		dp[i]=min
		if min==factor2{
			index2+=1
			factor2=2*dp[index2]
		}
		if min==factor3{
			index3+=1
			factor3=3*dp[index3]
		}
		if min==factor5{
			index5+=1
			factor5=5*dp[index5]
		}
	}
	return int(dp[n-1])
}