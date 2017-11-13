package dynamicprogramming

/*
You are given a list of non-negative integers, a1, a2, ..., an, and a target, S. Now you have 2 symbols + and -. For each integer, you should choose one from + and - as its new symbol.

Find out how many ways to assign symbols to make sum of integers equal to target S.

Example 1:
Input: nums is [1, 1, 1, 1, 1], S is 3.
Output: 5
Explanation:

-1+1+1+1+1 = 3
+1-1+1+1+1 = 3
+1+1-1+1+1 = 3
+1+1+1-1+1 = 3
+1+1+1+1-1 = 3

There are 5 ways to assign symbols to make the sum of nums be target 3.
//*/
//
//func FindTargetSumWays(nums []int, S int) int {
//     var dp [20][1000]int
//	 dp[0][0] = 1
//	for i:=0;i<len(nums);i++{
//		for _,a:=range dp{
//			dp[i+1][]
//		}
//	}
//}