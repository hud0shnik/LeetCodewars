package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	overval := 0
	l3 := &ListNode{
		Val:  0,
		Next: nil,
	}
	result := l3
	for l1 != nil || l2 != nil {
		l1Val := 0
		l2Val := 0
		if l1 != nil {
			l1Val = l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			l2Val = l2.Val
			l2 = l2.Next
		}
		sum := l1Val + l2Val + overval
		overval = sum / 10
		l3.Val = sum % 10

		if l1 == nil && l2 == nil && overval > 0 {
			l3.Next = &ListNode{
				Val:  overval,
				Next: nil,
			}
		}
		if l1 != nil || l2 != nil {
			l3.Next = &ListNode{
				Val:  0,
				Next: nil,
			}
			l3 = l3.Next
		}
	}
	return result
}

func main() {

}
