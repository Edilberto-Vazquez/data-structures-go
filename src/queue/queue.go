package queue

import "fmt"

type node struct {
	value string
	next  *node
}

type queue struct {
	first  *node
	last   *node
	length int
}

func NewQueue() *queue {
	return &queue{}
}

func (q *queue) Peek() *node {
	return q.first
}

func (q *queue) Enqueue(value string) {
	var newNode *node = &node{value: value}
	if q.length == 0 {
		q.first = newNode
		q.last = newNode
	} else {
		q.last.next = newNode
		q.last = newNode
	}
	q.length++
}

func (q *queue) Dequeue() {
	switch {
	case q.length == 0:
		fmt.Println("The queue is empty")
	case q.length == 1:
		q.first = nil
		q.last = nil
		q.length--
	default:
		q.first = q.first.next
		q.length--
	}
}
