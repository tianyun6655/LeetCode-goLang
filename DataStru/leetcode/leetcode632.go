package leetcode

import "fmt"

/**
You have k lists of sorted integers in ascending order. Find the smallest range that includes at least one number from each of the k lists.

We define the range [a,b] is smaller than range [c,d] if b-a < d-c or a < c if b-a == d-c.

Example 1:
Input:[[4,10,15,24,26], [0,9,12,20], [5,18,22,30]]
Output: [20,24]
Explanation:
List 1: [4, 10, 15, 24,26], 24 is in range [20,24].
List 2: [0, 9, 12, 20], 20 is in range [20,24].
List 3: [5, 18, 22, 30], 22 is in range [20,24].
 */


func SmallestRange(nums [][]int) []int {
	 leftNumber := nums[0][len(nums[0])-1]
	for _,a:=range nums{
		if leftNumber >= a[len(a)-1]{
			leftNumber = a[len(a)-1]
		}
	}
    rightNumber :=nums[0][0]
	for _,b:=range nums{
		temp:=b[0]
		for _,c :=range b{
			if c >leftNumber{
				temp= c
				break
			}
		}
		fmt.Println("temp",temp)

		if rightNumber <= temp{

			rightNumber = temp
		}
	}
	if leftNumber>rightNumber{
		leftNumber,rightNumber=rightNumber,leftNumber
	}
	result :=[]int{leftNumber,rightNumber}

	return result

}