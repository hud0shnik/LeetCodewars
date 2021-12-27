package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func hasCycle(head *ListNode) bool {
	fastScan, slowScan := head, head

	for {
		if fastScan == nil || fastScan.Next == nil {
			return false
		}
		fastScan = fastScan.Next.Next
		slowScan = slowScan.Next

		if fastScan == slowScan {
			return true
		}
	}
}

func main() {
}
