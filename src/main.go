package main

import (
	"data-structures-go/src/doublylinkedlist"
	"fmt"
)

func main() {
	// Sinly Linked List
	// ssl := singlylinkedlist.NewSinglylinkedlist()
	// ssl.Append(0)
	// ssl.Append(1)
	// ssl.Append(2)
	// ssl.Append(3)
	// ssl.Insert(4, 4)
	// ssl.Delete(20)
	// ssl.ShowValues()

	// Doubly Linked List
	dll := doublylinkedlist.NewDoublyLinkedList()
	dll.Append(0)
	dll.Append(1)
	dll.Append(2)
	dll.Append(3)
	dll.Append(4)
	// dll.Append(5)
	// dll.Prepend(-1)
	// dll.Insert(2, 50)
	// dll.Insert(6, 6)
	// dll.PrintList()
	// dll.Delete(3)
	dll.PrintList()
	fmt.Println(dll.GetIndex(4))
}
