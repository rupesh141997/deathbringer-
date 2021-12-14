package main

import "fmt"

type Node struct {
	val  int
	next *Node
}

type LinkedList struct {
	head   *Node
	length int
}

func main() {
	var head2 *Node
	for i := 1; i < 7; i++ {
		head2 = head2.insertAtEnd(i)
	}
	head2.printList()
}

// func (head Node) insertAtBegin(value int) *Node {
// 	var temp Node
// 	temp.val = value
// 	temp.next = &head
// 	return temp
// }

func (head *Node) insertAtEnd(value int) *Node {

	temp := &Node{}
	temp.val = value
	if head != nil {
		front := head
		for head.next != nil {
			head = head.next
		}
		head.next = temp
		return front
	} else {
		return temp
	}
}

func (head *Node) printList() {
	for head != nil {
		fmt.Println("Value:", head.val)
		head = head.next
	}
}

func (head *Node) printList2() {
	for head.next != nil {
		fmt.Println("val1", head.val)
		head = head.next
	}
	fmt.Println("val1", head.val)
}

func (l *LinkedList) construct(val int) {
	n := Node{}
	n.val = val
	if l.length == 0 {
		l.head = &n
		l.length++
		return
	}
	ptr := l.head
	for i := 0; i < l.length; i++ {
		if ptr.next == nil {
			ptr.next = &n
			l.length++
			return
		}
		ptr = ptr.next
	}
}

func ReverseLinkedList(head *Node) {
	if head == nil {
		return
	}
	ReverseLinkedList(head.next)
	fmt.Printf("%d ", head.val)
}
