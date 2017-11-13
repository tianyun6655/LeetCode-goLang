package dynamicprogramming

import (
	"testing"
	"fmt"
)

func TestCoinChange(t *testing.T) {
	testdata:=[]int{2}
	fmt.Println(CoinChange(testdata,3))
}
