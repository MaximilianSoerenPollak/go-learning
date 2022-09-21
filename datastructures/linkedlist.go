package main

import (
	"fmt"
)

type node struct {
	data int
	next *node
}

type linkedlist struct {
	head   *node
	length int
}

func (l *linkedlist) prepend(n *node) {
	second := l.head
	l.head = n
	l.head.next = second
	l.length++
}

func (l linkedlist) print_data() {
	toPrint := l.head
	for l.length != 0 {
		fmt.Printf("%d\t", toPrint.data)
		toPrint = toPrint.next
		l.length--
	}
}

func main() {
	mylist := linkedlist{}
	node1 := &node{data: 48}
	node2 := &node{data: 50}
	node3 := &node{data: 52}
	mylist.prepend(node1)
	mylist.prepend(node2)
	mylist.prepend(node3)
	mylist.print_data()
}
