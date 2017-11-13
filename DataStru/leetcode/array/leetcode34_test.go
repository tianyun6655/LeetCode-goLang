package array

import (
	"testing"
	"fmt"
)

func TestSearchRange(t *testing.T) {
	testdata:=[]int{1}
	target:=0

	fmt.Println(SearchRange(testdata,target))
}


func TestSearchRange2(t *testing.T) {
	testdata2:=[]int{5, 7, 7, 8, 8, 10}
	target2:=8

	fmt.Println(SearchRange2(testdata2,target2))
}