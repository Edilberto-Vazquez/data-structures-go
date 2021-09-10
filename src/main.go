package main

import (
	"fmt"

	"github.com/Edilberto-Vazquez/data-structures-go/src/binarysearchtree"
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
	// dll := doublylinkedlist.NewDoublyLinkedList()
	// dll.Append(0)
	// dll.Append(1)
	// dll.Prepend(-1)
	// dll.Insert(6, 6)
	// dll.Delete(3)
	// dll.PrintList()
	// fmt.Println(dll.GetIndex(4))

	// Stack
	// s := stack.NewStack()
	// s.Push("Tumtum")
	// s.Push("Pilingun")
	// s.Push("Edilberto")
	// s.Push("Maria")
	// s.Push("Ana")
	// fmt.Println(s.Peek())
	// s.Pop()
	// value, err := s.GetIndex(4)
	// if err != nil {
	// 	fmt.Println(err)
	// } else {
	// 	fmt.Println(value)
	// 	fmt.Println(s.LastItem())
	// }

	// Queue
	// q := queue.NewQueue()
	// q.Enqueue("Edilberto")
	// q.Enqueue("Maria")
	// q.Enqueue("Ana")
	// q.Peek()
	// fmt.Println(q.Peek())
	// q.Dequeue()
	// q.Dequeue()
	// q.Dequeue()

	// Binary Search Tree
	b := binarysearchtree.NewBinarySearchTree()
	b.Insert(10)
	b.Insert(4)
	b.Insert(20)
	b.Insert(2)
	b.Insert(8)
	b.Insert(17)
	b.Insert(170)
	// b.Delete(20)
	fmt.Println(b.Search(10))
	fmt.Println(b.Search(4))
	fmt.Println(b.Search(20))
	fmt.Println(b.Search(2))
	fmt.Println(b.Search(8))
	fmt.Println(b.Search(17))
	fmt.Println(b.Search(170))
	fmt.Println(b.Search(20))
	fmt.Println(b.Search(10))
	b.Delete(170)
	fmt.Println(b.Search(20))

}
