package main

type Node struct {
	Val   int
	Prev  *Node
	Next  *Node
	Child *Node
}

func flatten(root *Node) *Node {
	head := new(Node)
	curr := head
	var stack []*Node
	for root != nil || len(stack) > 0 {
		for root != nil {
			curr.Next = new(Node)
			*curr.Next = *root
			curr.Prev = curr
			curr = curr.Next
			stack = append(stack, root)
			root = root.Child
		}
		root = stack[len(stack)-1].Next
		stack = stack[:len(stack)-1]
	}
	return head.Next
}

//leetcode submit region end(Prohibit modification and deletion)
