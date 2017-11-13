package array

/*
Given an array with n objects colored red, white or blue, sort them so that objects of the same color are adjacent, with the colors in the order red, white and blue.

Here, we will use the integers 0, 1, and 2 to represent the color red, white, and blue respectively.

Note:
You are not suppose to use the library's sort function for this problem.
*/


func SortColors(nums []int)  {
   colorNumber:=make([]int,3)

	for _,color:=range nums{
		colorNumber[color]++
	}

	startIndex:=0
	endIndex:=-1
	for i:=0;i<len(colorNumber);i++{
		startIndex= endIndex+1
		endIndex = startIndex+colorNumber[i]-1
		for j:=startIndex;j<=endIndex;j++{
			nums[j]=i
		}
	}
}

func SortColors2(nums []int)  {
	zero:=-1 //前闭 和 后闭
	two:= len(nums)

	for i:=0;i<two;{
		if nums[i]==1{
			i++
		}else if nums[i]==2{
			two --
			nums[two],nums[i] = nums[i],nums[two]

		}else if nums[i]==0{
			nums[i],nums[zero+1] = nums[zero+1],nums[i]
			 zero++
			 i++

		}else{
			return
		}

	}

}