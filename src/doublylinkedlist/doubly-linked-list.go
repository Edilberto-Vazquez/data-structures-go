package doublylinkedlist

import "fmt"

type node struct {
	Value int
	Next  *node
	Prev  *node
}

type doublyLinkedList struct {
	Head   *node
	Tail   *node
	length int
}

func NewDoublyLinkedList() *doublyLinkedList {
	return &doublyLinkedList{}
}

func (d *doublyLinkedList) Length() int {
	return d.length
}

func (d *doublyLinkedList) GetIndex(index int) (currentNode *node) {
	switch {
	case index == 0:
		return d.Head
	case index >= d.length:
		fmt.Println("the index", index, "does not exist")
		return
	default:
		currentNode = d.Head
		counter := 0
		for counter != index {
			currentNode = currentNode.Next
			counter++
		}
		return
	}
}

func (d *doublyLinkedList) PrintList() {
	currentNode := d.Head
	counter := 0
	for counter < d.length {
		fmt.Println(currentNode, "->")
		currentNode = currentNode.Next
		counter++
	}
}

func (d *doublyLinkedList) Append(value int) {
	var newNode *node = &node{Value: value}
	if d.Head == nil {
		d.Head = newNode
		d.Tail = d.Head
	} else {
		newNode.Prev = d.Tail
		d.Tail.Next = newNode
		d.Tail = newNode
	}
	d.length++
}

func (d *doublyLinkedList) Prepend(value int) {
	var newNode *node = &node{Value: value}
	if d.Head == nil {
		d.Head = newNode
		d.Tail = d.Head
	} else {
		newNode.Next = d.Head
		d.Head.Prev = newNode
		d.Head = newNode
	}
	d.length++
}

func getIndex(index int, Head *node) (currentNode *node) {
	currentNode = Head
	counter := 0
	for counter != index-1 {
		currentNode = currentNode.Next
		counter++
	}
	return
}

func (d *doublyLinkedList) Insert(index int, value int) {
	switch {
	case index == 0:
		d.Prepend(value)
	case index < 0:
		fmt.Println("negative indeces are not allowed")
	case index >= d.length:
		d.Append(value)
	default:
		var newNode *node = &node{Value: value}
		firstPointer := getIndex(index, d.Head)
		holdingPointer := firstPointer.Next
		holdingPointer.Prev = newNode
		firstPointer.Next = newNode
		newNode.Next = holdingPointer
		newNode.Prev = firstPointer
		d.length++
	}
}

func (d *doublyLinkedList) Delete(index int) {
	switch {
	case index == 0:
		d.Head = d.Head.Next
		d.Head.Prev = nil
		d.length--
	case index < 0:
		fmt.Println("negative indeces are not allowed")
	case index > d.length:
		fmt.Println("the index", index, "does not exist")
	default:
		firsPointer := getIndex(index, d.Head)
		holdingPointer := firsPointer.Next
		firsPointer.Next = holdingPointer.Next
		if firsPointer.Next == nil {
			d.Tail = firsPointer
		} else {
			holdingPointer.Next.Prev = firsPointer
		}
		d.length--
	}
}
