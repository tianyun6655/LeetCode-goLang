package leetcode

import (
	"fmt"
)

/**


Given a triangle, find the minimum path sum from top to bottom. Each step you may move to adjacent numbers on the row below.

For example, given the following triangle
[
     [2],
    [3,4],
   [6,5,7],
  [4,1,8,3]
]
The minimum path sum from top to bottom is 11 (i.e., 2 + 3 + 5 + 1 = 11).

Note:
Bonus point if you are able to do this using only O(n) extra space, where n is the total number of rows in the triangle.


 */

func MinimumTotal(triangle [][]int) int {
	result :=make([]int,1)
	result[0] = triangle[0][0]
	indexA :=make([]int,1)
	indexA[0]=0
	for i:=1;i<len(triangle);i++{
		temp :=make([]int,len(result)*2)
		tempn:=0
		tempIndex :=make([]int,len(indexA)*2)
		tempi:=0
		for index:=0;index<len(indexA);index++{
			tempIndex[tempi] = indexA[index]
			tempi++
			tempIndex[tempi] = indexA[index]+1
			tempi++
		}
		k:=0
		fmt.Println(tempIndex)
		for j:=0;j<len(result);j++{
			temp[tempn]=result[j]+triangle[i][tempIndex[k]]
			tempn++

			k++
			temp[tempn] = result[j]+triangle[i][tempIndex[k]]
			k++
			tempn++

		}

		result  = temp
		indexA = tempIndex

	}
	min:=result[0]

	fmt.Println(result)

	for _,a :=range result{
		if a<min{
			min = a
		}
	}
	return min
}

