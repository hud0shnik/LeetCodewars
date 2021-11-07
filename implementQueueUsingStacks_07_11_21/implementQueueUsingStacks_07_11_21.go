package main

type (
	Node struct {
		val  int
		next *Node
	}

	MyQueue struct {
		head *Node
		tail *Node
	}
)

func Constructor() MyQueue {
	return MyQueue{}
}

func (this *MyQueue) Push(x int) {
	element := &Node{val: x}
	if this.tail == nil {
		this.tail, this.head = element, element
	} else {
		this.tail = element
		this.tail.next = element
	}
}

func (this *MyQueue) Pop() int {
	result := this.head
	if result != nil {
		this.head = this.head.next
		if this.head == nil {
			this.tail = nil
		}
		return result.val
	}
	return 0
}

func (this *MyQueue) Peek() int {
	if this.head != nil {
		return this.head.val
	}
	return 0
}

func (this *MyQueue) Empty() bool {
	if this.head == nil {
		return true
	}
	return false
}
