package main

type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Next  *Node
}

func connect(root *Node) *Node {
	if root == nil {
		return nil
	}

	q := []*Node{root}

	for len(q) > 0 {
		size := len(q)
		for i := 0; i < size; i++ {
			if q[0].Left != nil {
				q = append(q, q[0].Left)
			}
			if q[0].Right != nil {
				q = append(q, q[0].Right)
			}

			if i != size-1 {
				q[0].Next = q[1]
			} else {
				q[0].Next = nil
			}
			q = q[1:]
		}
	}

	return root
}

func main() {
}
