package main

import (
	"data-structures-go/src/singlylinkedlist"
)

func main() {
	ssl := singlylinkedlist.NewSinglylinkedlist()
	ssl.Append(0)
	ssl.Append(1)
	ssl.Append(2)
	ssl.Append(3)
	ssl.Insert(4, 4)
	ssl.Delete(20)
	ssl.ShowValues()
}
