/*
 @Author : dongyang
 @Time   : 2021/9/26 10
 @Desc   : 链表
*/
package code

import "fmt"

// 合并多级链表
type Node struct {
	Val   int
	Prev  *Node
	Next  *Node
	Child *Node
}

func flatten(root *Node) *Node {
	var head = new(Node)
	if root == nil {
		return nil
	}
	curr := head
	var stack []*Node
	for root != nil || len(stack) > 0 {
		for root != nil {
			curr.Next = &Node{
				Val:  root.Val,
				Prev: curr,
				Next: nil,
			}
			curr = curr.Next
			stack = append(stack, root)
			root = root.Child
		}
		root = stack[len(stack)-1].Next
		stack = stack[:len(stack)-1]

	}
	head.Next.Prev = nil
	return head.Next
}

//基本计算器
func calculate(s string) int {
	var (
		stack []int
		num   = 0
		ops   = byte('+')
		bs    = []byte(s)
		ans   int
	)
	for i := 0; i < len(bs); i++ {
		isnum := '0' <= bs[i] && bs[i] <= '9'
		if isnum {
			num = num*10 + int(bs[i]-'0')
		}
		if !isnum && bs[i] != ' ' || i == len(bs)-1 {
			switch ops {
			case '+':
				stack = append(stack, num)
			case '-':
				stack = append(stack, -num)
			case '*':
				stack[len(stack)-1] *= num
			case '/':
				stack[len(stack)-1] /= num
			default:
				return 0
			}
			ops = bs[i]
			num = 0
		}
	}
	for i := 0; i < len(stack); i++ {
		ans += stack[i]
	}
	return ans
}

type ListNode struct {
	Val  int
	Next *ListNode
}

//19. 删除链表的倒数第 N 个结点
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	a, b, c := &ListNode{Val: 0, Next: head}, head, head.Next
	now:=a
	for c != nil {
		c = c.Next
		if n > 0 {
			n--
		}
		if n == 0 {
			fmt.Println(b)
			a = b
			b = b.Next
		}
	}
	fmt.Println(b)
	a.Next = b.Next
	return now.Next
}
