package sort

import (
	"testing"
	"fmt"
	"time"
	"math/rand"
)

func TestQuickSort(t *testing.T) {
	testdata:=make([]int,100000)
	random:=rand.New(rand.NewSource(time.Now().UnixNano()))

	for i:=0;i<len(testdata);i++{
		testdata[i]= random.Int()%10
	}

	startTime:= time.Now()

	QuickSort2(testdata,len(testdata))

	elapsed := time.Since(startTime)

	fmt.Println("elapsed: ",elapsed)

	startTime2:= time.Now()

	QuickSort(testdata,len(testdata))

	elapsed2 := time.Since(startTime2)

	fmt.Println("elapsed2: ",elapsed2)


}
