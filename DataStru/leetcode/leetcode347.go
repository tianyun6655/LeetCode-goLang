package leetcode

func TopKFrequent(nums []int, k int) []int {
	 freqMap:=make(map[int]int)

	for _,a:=range nums{
		freqMap[a]++
	}

	heap:=NewMaxHeap(len(nums))

	for key,value:=range freqMap{
		elementI :=&element{frequence:value,number:key}
		if elementI.frequence >heap.Extract().frequence{
			heap.Insert(*elementI)

		}
	}

	result :=make([]int,k)

	for i:=0;i<len(result);i++{
		result[i]=heap.Extract().number
	}
   return result
}
type element struct {
   frequence int
	number int
}

type MaxHeap struct {
	data []element
	capacity int
	count int
}


func NewMaxHeap(capacity int) *MaxHeap{

	//start from index 1
	data :=make([] element,capacity+1)

	return &MaxHeap{
		data:data,
		count:0,
		capacity:capacity,
	}
}


func(heap *MaxHeap)Size() int{
	return heap.count
}


func(heap *MaxHeap)Insert(item element){
	heap.count++
	if heap.count<=heap.capacity {
		heap.data[heap.count] = item
		heap.shifUp(heap.count)
	}
}

func(heap *MaxHeap)shifUp(index int){
	for heap.data[index/2].frequence<heap.data[index].frequence&& index>1{
		heap.data[index/2],heap.data[index] = heap.data[index],heap.data[index/2]
		index = index/2
	}
}

func (heap *MaxHeap)Extract() element{

	item :=heap.data[1]

	heap.data[heap.count],heap.data[1]=heap.data[1],heap.data[heap.count]

	heap.count--

	heap.shiftdown(1)
	return item
}

func(heap *MaxHeap)shiftdown(index int){
	for 2*index<= heap.count{
		j:=2*index
		if (j+1<=heap.count&&heap.data[j+1].frequence>heap.data[j].frequence){
			j++

		}
		if heap.data[index].frequence>heap.data[j].frequence{
			break
		}
		heap.data[index],heap.data[j]=heap.data[j],heap.data[index]

		index = j

	}
}