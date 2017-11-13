package sort

import (
	"math/rand"
	"time"
)
var randSelector *rand.Rand

func QuickSort(arr []int,n int){

	randSelector=rand.New(rand.NewSource(time.Now().UnixNano()))

   quickSortHelper(arr,0,n-1)



}


func quickSortHelper(arr []int, left,right int){

	if left>=right{
		return
	}


	basePoint:=partition(arr,left,right)

	quickSortHelper(arr,left,basePoint-1)
	quickSortHelper(arr,basePoint+1,right)

}

func partition(arr[]int,left,right int)int{

	randNumber:=randSelector.Int()%(right-left+1)+left

	arr[randNumber],arr[left]=arr[left],arr[randNumber]

	basePoint:=arr[left]

	j:=left

	for i:=j+1;i<=right;i++{
		if arr[i]<basePoint{
			arr[j+1],arr[i] = arr[i],arr[j+1]
			j++
		}
	}

	arr[j],arr[left] = arr[left],arr[j]

	return j
}


func QuickSort2(arr []int,n int){

	randSelector=rand.New(rand.NewSource(time.Now().UnixNano()))

	quickSortHelper2(arr,0,n-1)



}
func quickSortHelper2(arr []int, left,right int){

	if left>=right{
		return
	}


	basePoint:=partition2(arr,left,right)

	quickSortHelper2(arr,left,basePoint-1)
	quickSortHelper2(arr,basePoint+1,right)

}

func partition2(arr[]int,left,right int)int{

	randNumber:=randSelector.Int()%(right-left+1)+left

	arr[randNumber],arr[left]=arr[left],arr[randNumber]
	base:=arr[left]

	i:=left+1
	j:=right
	for true{
		for i<=right&&arr[i]<base{
			i++
		}
		for j>=left&&arr[j]>base{
			j--
		}

		if i >j{
			break
		}
		arr[i],arr[j]=arr[j],arr[i]
		i++
		j--

	}
	arr[left],arr[j] =arr[j],arr[left]

	return j
}
