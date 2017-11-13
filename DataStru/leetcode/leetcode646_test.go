package leetcode

import (
	"testing"
	"fmt"
)

func TestFindLongestChain(t *testing.T) {
	testdata:=[][]int{{1,2},{2,3},{3,4}}
	retsult :=FindLongestChain(testdata)
	fmt.Println(retsult)
}
