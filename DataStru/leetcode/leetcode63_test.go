package leetcode

import (
	"testing"

	"fmt"
)

func TestUniquePathsWithObstacles(t *testing.T) {
	testdata :=[][]int{{0,0,0},{0,1,0},{0,0,0}}
	fmt.Println(UniquePathsWithObstacles(testdata))
}
