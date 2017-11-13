package tree

import (
	"fmt"
	"math"
)

type node struct{
	value int
	leftNode *node
	rightNode *node
}

type BinarySearchTree struct {
    rootNode *node
    size int
}

func NewBinarySearchTree() *BinarySearchTree{
	return &BinarySearchTree{size:0}
}


func(bst *BinarySearchTree)GetSize() int{
	return bst.size

}

func (bst *BinarySearchTree)IsEmpty() bool{
	return bst.size==0
}

func (bst *BinarySearchTree)Insert(value int){

    bst.rootNode = insert(bst.rootNode,value)

}

func (bst *BinarySearchTree)Contains(value int) bool{
   return contains(bst.rootNode,value)
}


func contains (rootNode *node ,value int) bool{

	if rootNode==nil{
		return false
	}
		if rootNode.value == value {
			return true
		} else if rootNode.value > value {
			return contains(rootNode.leftNode, value)
		} else  {
			return contains(rootNode.rightNode, value)
		}

}


//返回插入节点后对应的根
func insert(rootNode *node, value int) *node{

	if rootNode==nil{

		return &node{value:value}

	}

	if value == rootNode.value{
		rootNode.value = value
	}else if value < rootNode.value{
		rootNode.leftNode = insert(rootNode.leftNode,value)
	}else if value >rootNode.value{
		rootNode.rightNode = insert(rootNode.rightNode,value)
	}

	return rootNode

}
func (bst *BinarySearchTree)InsetrWithoutRe(value int){

	if bst.rootNode==nil{
		bst.rootNode = &node{value:value}
		return
	}

	 currentNode :=bst.rootNode
     parentNode :=bst.rootNode

	 for currentNode!=nil{

	 	parentNode = currentNode

		 if value < currentNode.value{
			 currentNode = currentNode.leftNode
		 }else if value >currentNode.value{
			 currentNode= currentNode.rightNode
		 }
	 }

   if parentNode.value >value{
   	parentNode.leftNode = &node{value:value}
   }else{
   	parentNode.rightNode = &node{value:value}
   }


}


//遍历：前序：先访问自身，在递归 访问左右，中序：先左，在自身最后右边。后：先左右后自身


func (bst *BinarySearchTree)PreOrder(){
	preOrder(bst.rootNode)
}

func preOrder(rootNode *node){


	if rootNode!=nil{
		fmt.Println(rootNode.value)
		preOrder(rootNode.leftNode)
		preOrder(rootNode.rightNode)
	}
	
}

func (bst *BinarySearchTree)InOrder(){
	inOrder(bst.rootNode)
}

func inOrder(rootNode *node){

	if rootNode!=nil{
		inOrder(rootNode.leftNode)
		fmt.Println(rootNode.value)
		inOrder(rootNode.rightNode)
	}
}

func (bst *BinarySearchTree)PostOrder(){
	postOrder(bst.rootNode)
}

func postOrder(rootNode *node){

	if rootNode!=nil{
		postOrder(rootNode.leftNode)
		postOrder(rootNode.rightNode)
		fmt.Println(rootNode.value)

	}
}