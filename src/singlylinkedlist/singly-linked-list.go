package singlylinkedlist

import "fmt"

type node struct {
	Value int
	Next  *node
}

type singlylinkedlist struct {
	Head   *node
	Tail   *node
	length int
}

func NewSinglylinkedlist() *singlylinkedlist {
	return &singlylinkedlist{}
}

func (sll *singlylinkedlist) Length() int {
	return sll.length
}

func (sll *singlylinkedlist) Append(value int) {
	var newNode *node = &node{Value: value}
	if sll.Head == nil {
		sll.Head = newNode
		sll.Tail = sll.Head
	} else {
		sll.Tail.Next = newNode
		sll.Tail = newNode
	}
	sll.length++
}

func (sll *singlylinkedlist) Prepend(value int) {
	var newNode *node = &node{Value: value}
	if sll.Head == nil {
		sll.Head = newNode
		sll.Tail = sll.Head
	} else {
		newNode.Next = sll.Head
		sll.Head = newNode
	}
	sll.length++
}

func getIndex(index int, head *node) (currentNode *node) {
	currentNode = head
	counter := 0
	for counter != index-1 {
		currentNode = currentNode.Next
		counter++
	}
	return
}

func (sll *singlylinkedlist) Insert(index int, value int) {
	switch {
	case index == 0:
		sll.Prepend(value)
	case index >= sll.length:
		sll.Append(value)
	default:
		var newNode *node = &node{Value: value}
		firsPointer := getIndex(index, sll.Head)
		holdingPointer := firsPointer.Next
		firsPointer.Next = newNode
		newNode.Next = holdingPointer
		sll.length++
	}
}

func (sll *singlylinkedlist) Delete(index int) {
	switch {
	case index == 0:
		sll.Head = sll.Head.Next
		sll.length--
	case index >= sll.length:
		// firsPointer := getIndex(sll.length-1, sll.Head)
		// holdingPointer := firsPointer.Next
		// firsPointer.Next = holdingPointer.Next
		fmt.Printf("the index %d does not exist\n", index)
	default:
		firsPointer := getIndex(index, sll.Head)
		holdingPointer := firsPointer.Next
		firsPointer.Next = holdingPointer.Next
		sll.length--
	}
}

func (sll *singlylinkedlist) ShowValues() {
	counter := 0
	currentNode := sll.Head
	for counter != sll.length-1 {
		fmt.Println(currentNode, "->", currentNode.Next)
		currentNode = currentNode.Next
		counter++
	}
}
