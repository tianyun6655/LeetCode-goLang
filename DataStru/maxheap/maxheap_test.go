package maxheap

import (
	"testing"
	"fmt"
)

func TestNewMaxHeap(t *testing.T) {
	heap:=NewMaxHeap(9)
	heap.Insert(62)
	heap.Insert(30)
	heap.Insert(22)
	heap.Insert(16)
	heap.Insert(41)
    heap.Insert(28)
	fmt.Println(heap)
	fmt.Println("===============")
	heap.Extract()
	heap.Extract()

	fmt.Println(heap)
}
