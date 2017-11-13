package array

import "math"

/*
Find the contiguous subarray within an array (containing at least one number) which has the largest product.

For example, given the array [2,3,-2,4],
the contiguous subarray [2,3] has the largest product = 6.

*/

//暴力解  N^2
func maxProduct(nums []int) int {
	max:=math.MinInt64

	if len(nums)==1{
		return nums[0]
	}
	for i:=0;i<len(nums);i++{
		  preResult:=nums[i]

		if preResult>max{
			max = preResult
		}
		for j:=i+1;j<len(nums);j++{
             preResult*=nums[j]

			if preResult>max{
				max = preResult
			}
		}

	}

	return max
}

//DP
func maxProduct2(nums []int) int {
	maxValue:=nums[0]
	minValue:=nums[0]
	ansValue:=nums[0]

	for i:=1;i<len(nums);i++{
		currentMax:= maxValue*nums[0]
		currentMin := minValue*nums[0]

		maxValue = max(max(currentMax,currentMin),nums[i])
		minValue = min(max(currentMax,currentMin),nums[i])
		ansValue = max(maxValue,minValue)
	}

	return ansValue
}

func max(a,b int) int{
	if a>b{
		return a
	}else{
		return b
	}
}

func min(a,b int) int{
	if a<b{
		return a
	}else{
		return b
	}
}