package binarysearchtree

type node struct {
	left  *node
	rigth *node
	value int
}

type binarySearchTree struct {
	root *node
}

func NewBinarySearchTree() *binarySearchTree {
	return &binarySearchTree{}
}

func (b *binarySearchTree) Insert(value int) {
	var newNode *node = &node{value: value}
	if b.root == nil {
		b.root = newNode
	} else {
		var currentNode *node = b.root
		for {
			if value < currentNode.value {
				if currentNode.left == nil {
					currentNode.left = newNode
					break
				}
				currentNode = currentNode.left
			} else {
				if currentNode.rigth == nil {
					currentNode.rigth = newNode
					break
				}
				currentNode = currentNode.rigth
			}
		}
	}
}

func (b binarySearchTree) Search(value int) (firstNode *node) {
	for {
		switch {
		case b.root == nil:
			return nil
		case value < b.root.value:
			b.root = b.root.left
		case value > b.root.value:
			b.root = b.root.rigth
		default:
			return b.root
		}
	}
}

func (b *binarySearchTree) Delete(value int) {
	for {
		if b.root == nil {
			break
		} else if value < b.root.value {
			if value == b.root.left.value {
				b.root.left = nil
				break
			} else {
				b.root = b.root.left
			}
		} else if value > b.root.value {
			if value == b.root.rigth.value {
				b.root.rigth = nil
				break
			} else {
				b.root = b.root.rigth
			}
		} else {
			b.root = nil
			break
		}
	}
}
