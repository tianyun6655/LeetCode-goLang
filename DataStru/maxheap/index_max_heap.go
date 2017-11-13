package maxheap
type IndexMaxHeap struct {
	data []int
	index []int
	capacity int
	count int
}

//
//func NewIndexMaxHeap(capacity int) *IndexMaxHeap{
//
//	//start from index 1
//	data :=make([] int,capacity+1)
//    index :=make([]int, capacity+1)
//	return &IndexMaxHeap{
//		data:data,
//		index:index,
//		count:0,
//		capacity:capacity,
//	}
//}
//
//
//func(heap *MaxHeap)Size() int{
//	return heap.count
//}
//
//
//func(heap *IndexMaxHeap)Insert(item int){
//	heap.count++
//	if heap.count<=heap.capacity {
//		heap.data[heap.count] = item
//		heap.shifUp(heap.count)
//	}
//}
//
//func(heap *MaxHeap)shifUp(index int){
//	for heap.data[index/2]<heap.data[index]&& index>1{
//		heap.data[index/2],heap.data[index] = heap.data[index],heap.data[index/2]
//		index = index/2
//	}
//}
//
//func (heap *MaxHeap)Extract() int{
//	if heap.count<=0{
//		return 0
//	}
//	item :=heap.data[1]
//
//	heap.data[heap.count],heap.data[1]=heap.data[1],heap.data[heap.count]
//
//	heap.count--
//
//	heap.shiftdown(1)
//	return item
//}


//func(heap *MaxHeap)shiftdown(index int){
//	for 2*index<= heap.count{
//		j:=2*index
//		if (j+1<=heap.count&&heap.data[j+1]>heap.data[j]){
//			j++
//
//		}
//		if heap.data[index]>heap.data[j]{
//			break
//		}
//		heap.data[index],heap.data[j]=heap.data[j],heap.data[index]
//
//		index = j
//
//	}
//}