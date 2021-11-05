package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeElements(head *ListNode, val int) *ListNode {
	current := head

	for {
		if current.Next.Val == val {
			current.Next = current.Next.Next
		} else {
			current = current.Next
		}
		if current.Next == nil {
			break
		}
	}

	if head.Val == val {
		return head.Next
	}

	return head
}

func main() {
}
