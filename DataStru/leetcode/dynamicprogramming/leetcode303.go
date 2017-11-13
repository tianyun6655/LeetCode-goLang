package dynamicprogramming
/*
Given an integer array nums, find the sum of the elements between indices i and j (i ≤ j), inclusive.

Example:
Given nums = [-2, 0, 3, -5, 2, -1]

sumRange(0, 2) -> 1
sumRange(2, 5) -> -1
sumRange(0, 5) -> -3
*/


type NumArray struct {
	nums []int

}


func Constructor(nums []int) NumArray {
	sumNums:=make([]int,len(nums))
	for i:=1;i<len(nums);i++{
		sumNums[i]+=sumNums[i-1]
	}
	return NumArray{
		nums:sumNums,
	}
}


func (this *NumArray) SumRange(i int, j int) int {

	if i==0{
		return this.nums[j]
	}
	return this.nums[j] - this.nums[i-1]

}
