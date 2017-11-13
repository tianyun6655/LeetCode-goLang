package leetcode

import (
	"testing"
	"fmt"
)

func TestMinimumTotal(t *testing.T) {
	testdata:=[][]int{{2},{3,4},{6,5,9},{4,4,8,0}}
	fmt.Println(MinimumTotal(testdata))
}
