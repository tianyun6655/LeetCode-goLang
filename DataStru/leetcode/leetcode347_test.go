package leetcode

import (
	"testing"
	"fmt"
)

func TestTopKFrequent(t *testing.T) {
	arr:=[]int{1,1,1,2,2,3}
	result :=TopKFrequent(arr,2)
	fmt.Println(result)
}
