package tree

import (
	"testing"
	"fmt"
)

func TestBinarySearchRecurison(t *testing.T) {

	testdata:=[]int{2,3,4,5,6,7,8}

	target:=7

	fmt.Println(BinarySearchRecurison(testdata,len(testdata),target))

}

