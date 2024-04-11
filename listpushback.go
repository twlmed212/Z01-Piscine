package piscine

import "fmt"

type node struct {
	data int
	next *node
}
type LinkedList struct {
	head *node
}

func (ll *LinkedList) insert(data int) {
	newNode := &node{data: data}
	newNode.next = ll.head
	ll.head = newNode
}
func (ll *LinkedList) Print() {
	temp := ll.head
	for temp != nil {
		fmt.Printf("%d -> ", temp.data)
		temp = temp.next
	}
	fmt.Println("nil")
}
