package array

import (
	"testing"
	"fmt"
)

func TestMoveZeroes(t *testing.T) {
	testdata:=[]int{0,1,0,3,12}

	MoveZeroes(testdata)

	fmt.Println(testdata)

}
