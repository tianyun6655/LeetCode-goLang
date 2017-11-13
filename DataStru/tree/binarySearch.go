package tree


//作用在有序的数列中，1946年提出



func binarySearch(arr []int, n int, target int) int{

	l :=0
	r:=len(arr)-1

	for l<=r{
		//有bug
        // mid:=(l+r)>>1
        mid:=l+(r-l)/2

        if arr[mid]==target{
        	return mid
		}
		if target<arr[mid]{
			r=mid-1
		}else{
			l = mid+1
		}
	}

	return -1
}


func BinarySearchRecurison(arr []int, n int,target int) int{

   return  binarySearhHelper(arr,0,len(arr)-1,target)

}

func binarySearhHelper(arr []int,left,right,target int) int{

	if left < right {
        mid:=left+(right-left)/2
        if arr[mid]==target{
        	return mid
		}
	if arr[mid]>target{
		return  binarySearhHelper(arr,left,mid-1,target)
	   }else {
	     return binarySearhHelper(arr,mid+1,right,target)
	}

	}else{
		return -1
	}
}

