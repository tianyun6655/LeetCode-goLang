package dynamicprogramming

import (
	"testing"
	"fmt"
)

func TestCanPartition(t *testing.T) {
	testdata:=[]int{1,2,5}
	fmt.Println(CanPartition(testdata))
}
