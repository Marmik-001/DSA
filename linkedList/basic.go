package linkedlist

import "fmt"

type node struct {
	next *node
	data int
}

type linkedlist struct {
	head   *node
	tail *node
	length int
}

func (l *linkedlist) prepend(value int) {
	newNode := node{data: value}
	if l.head != nil {
		newNode.next = l.head
		l.head = &newNode
		l.length++
	} else {
		l.head = &newNode
		l.tail = &newNode
		l.length++
	}
	// fmt.Println(l.head.data ,l.tail.data)

}

func (l *linkedlist) append(value int) {
	newNode := node{data:value}
	if l.tail != nil {
		l.tail.next = &newNode
		l.tail = &newNode
		l.length++
	}else {
		l.head= &newNode
		l.tail = &newNode
		l.length++
	}
	return
}

func (l *linkedlist) getValues() {

	if l.head == nil {
		return
	}

	currentNode := l.head
	for currentNode != nil {
		fmt.Println(currentNode.data)
		currentNode = currentNode.next
	}

	// for i := 0; i < l.length; i++ {
	// 	fmt.Printf("The current node:%v \n the next node address: %v", l.head.data, l.head.next)
	// 	nextNode := l.head.next
	// 	fmt.Printf("%v current number... The current node:%v \n the next node address: %v ", i, nextNode.data, nextNode.next)
	// }
}

func (l *linkedlist) deleteNodeWithValue(value int) {
	if l.length == 0 {
		return
	}

	if l.head.data == value {
		l.head = l.head.next
		l.length--
		return

	}

	prevNode := l.head
	currentNode := l.head.next
	for currentNode != nil {
		if currentNode.data == value {
			prevNode.next = currentNode.next
			l.length--
			return
		}
		prevNode = currentNode
		currentNode = currentNode.next
	}
	// prevToDelete := l.head
	// for prevToDelete.next.data != value {
	// 	if prevToDelete.next.next == nil {
	// 		return
	// 	}
	// 	prevToDelete = prevToDelete.next
	// }
	// prevToDelete.next = prevToDelete.next.next
	// l.length--
}

func (l *linkedlist) reverseLinkedList(){

	var prev *node = nil
	current := l.head

	for current != nil {
		next := current.next
		current.next = prev
		prev = current
		current = next
	}
	l.head = prev
}

func (l *linkedlist) findCycle() bool{
	slow, fast := l.head , l.head
	for fast != nil && fast.next != nil {
		slow = slow.next
		fast = fast.next.next
		if slow == fast {
			return true
		}
	}
	return false
}

func Basic() {
	// You can use linkedlist and its methods here
	myList := linkedlist{}

	myList.prepend(32)
	myList.prepend(44)
	myList.prepend(3)
	myList.prepend(66)

	// myList.deleteNodeWithValue(32)

	fmt.Println("old values")
	myList.getValues()

	myList.append(100)
	fmt.Println("new values")
	myList.reverseLinkedList()
	myList.getValues()
	cycle := myList.findCycle()
	fmt.Println(cycle)
}
