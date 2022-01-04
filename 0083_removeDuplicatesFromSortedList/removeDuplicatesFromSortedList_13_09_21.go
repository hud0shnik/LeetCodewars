package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	temp := head

	for temp.Next != nil {
		if temp.Val == temp.Next.Val {
			temp.Next = temp.Next.Next
		} else {
			temp = temp.Next
		}
	}

	return head
}

func main() {

}

/*


func deleteDuplicates(head *ListNode) *ListNode {
    if head == nil {
        return nil
    }
    if head.Next == nil {
        return head
    }
    if head.Val == head.Next.Val {
        return deleteDuplicates(head.Next)
    }
    head.Next = deleteDuplicates(head.Next)
    return head
}

*/
