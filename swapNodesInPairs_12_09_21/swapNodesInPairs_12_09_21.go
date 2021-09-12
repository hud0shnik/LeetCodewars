package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func swapPairs(head *ListNode) *ListNode {
	res := &ListNode{Next: head}
	tm := res
	for (tm != nil) && (tm.Next != nil) && (tm.Next.Next != nil) {
		tm, tm.Next, tm.Next.Next, tm.Next.Next.Next = tm.Next, tm.Next.Next, tm.Next.Next.Next, tm.Next
	}
	return res.Next
}

func swapPairs2(head *ListNode) *ListNode {
	dummy := &ListNode{0, head}
	prev := dummy
	for head != nil && head.Next != nil {
		first := head
		second := head.Next
		prev.Next = second
		first.Next = second.Next
		second.Next = first
		prev = first
		head = head.Next
	}

	return dummy.Next
}

func main() {

}
