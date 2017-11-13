package dynamicprogramming

import (
	"testing"
	"fmt"
)

func TestMaximalSquare(t *testing.T) {
   testData:=[][]byte{{1,0,1,0,0},{1,0,1,1,1},{1,1,1,1,1},{1,0,0,1,0}}
	fmt.Println(MaximalSquare(testData))
}
