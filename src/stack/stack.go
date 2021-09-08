package stack

import (
	"fmt"
)

type node struct {
	value string
	next  *node
}

type stack struct {
	top    *node
	bottom *node
	length int
}

func NewStack() *stack {
	return &stack{}
}

func (s *stack) Peek() *node {
	return s.top
}

func (s *stack) Push(value string) {
	var newNode *node = &node{value: value}
	if s.length == 0 {
		s.top = newNode
		s.bottom = newNode
	} else {
		var holdingPointer *node = s.top
		s.top = newNode
		s.top.next = holdingPointer
	}
	s.length++
}

func (s *stack) Pop() {
	if s.length <= 0 {
		s.top = nil
		s.bottom = nil
		s.length = 0
	} else {
		s.top = s.top.next
		s.length--
	}
}

func (s *stack) FirstItem() *node {
	return s.top
}

func (s *stack) LastItem() *node {
	return s.bottom
}

func (s *stack) GetIndex(index int) (*node, error) {
	var counter int = 0
	var currentNode *node = s.top
	switch {
	case index >= s.length:
		return nil, fmt.Errorf("you have exceeded the maximum number of indexes, index %v does not exist", index)
	case index < 0:
		return nil, fmt.Errorf("negative values are not allowed, the stack starts at index 0")
	default:
		for counter != index {
			currentNode = currentNode.next
			counter++
		}
		return currentNode, nil
	}
}
