package tree

import "testing"

func TestNewBinarySearchTree(t *testing.T) {
	bst :=NewBinarySearchTree()

	bst.Insert(10)
	bst.Insert(6)
	bst.Insert(13)
	bst.Insert(4)
	bst.Insert(7)
	bst.PreOrder()


	bst2 :=NewBinarySearchTree()
	bst2.InserWithoutRe(10)
	bst2.InserWithoutRe(6)
	bst2.InserWithoutRe(13)
	bst2.InserWithoutRe(4)
	bst2.InserWithoutRe(7)
	bst2.PreOrder()

}