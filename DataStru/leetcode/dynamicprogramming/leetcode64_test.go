package dynamicprogramming

import (
	"testing"
	"fmt"
)

func TestMinPathSum(t *testing.T) {
	testdata:=[][]int{{1,3,1},{1,5,1},{4,2,1}}
	fmt.Println(MinPathSum(testdata))
}
