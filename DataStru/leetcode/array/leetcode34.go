package array

import (
)

func SearchRange(nums []int, target int) []int {

	firstIndex:=-1


	for i:=0;i<len(nums);i++{
		if nums[i]==target{
			firstIndex = i
			break
		}
	}

	secondeIndex:=-1

	if firstIndex==-1{
		return []int{-1,-1}
	}

	for i:=firstIndex;i<len(nums);i++{
		if nums[i]==target{
			secondeIndex= i
		}
	}
	return []int{firstIndex,secondeIndex}
}


func SearchRange2(nums []int, target int) []int {


    firstIndex:=binarySearchFirst(nums,target)
    secondIndex:=binarySearchLast(nums,target)


    return []int{firstIndex,secondIndex}

    }

func binarySearchFirst(nums []int, target int) int{
	targetIndex:=-1
	start:=0
	end:=len(nums)-1

	for start<=end{
		mid:=(start+end)>>1

		if nums[mid]>=target{
			end =mid-1
		}else {
			start =mid+1
		}

		if nums[mid]==target{
			targetIndex=mid
		}
	}
	return targetIndex
}


func binarySearchLast(nums []int, target int) int{
	targetIndex:=-1
	start:=0
	end:=len(nums)-1

	for start<=end{
		mid:=(start+end)>>1

		if nums[mid]<=target{
			start =mid+1
		}else {
			end =mid-1
		}

		if nums[mid]==target{
			targetIndex=mid
		}
	}
	return targetIndex
}