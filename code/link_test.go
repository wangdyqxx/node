/*
 @Author : dongyang
 @Time   : 2021/9/26 15
 @Desc   :
*/
package code

import (
	"fmt"
	"testing"
)

func TestLink_flatten(t *testing.T) {
	var n1, n2, n3, n4, n5, n6, n7, n8, n9, n10, n11, n12 = new(Node), new(Node), new(Node), new(Node), new(Node), new(Node), new(Node), new(Node), new(Node), new(Node), new(Node), new(Node)

	n1.Val = 1
	n1.Prev = nil
	n1.Next = n2
	n1.Child = nil

	n2.Val = 2
	n2.Prev = n1
	n2.Next = n3
	n2.Child = nil

	n3.Val = 3
	n3.Prev = n2
	n3.Next = n4
	n3.Child = n7

	n4.Val = 4
	n4.Prev = n3
	n4.Next = n5
	n4.Child = nil

	n5.Val = 5
	n5.Prev = n4
	n5.Next = n6
	n5.Child = nil

	n6.Val = 6
	n6.Prev = n5
	n6.Next = nil
	n6.Child = nil

	n7.Val = 7
	n7.Prev = nil
	n7.Next = n8
	n7.Child = nil

	n8.Val = 8
	n8.Prev = n7
	n8.Next = n9
	n8.Child = n11

	n9.Val = 9
	n9.Prev = n8
	n9.Next = n10
	n9.Child = nil

	n10.Val = 10
	n10.Prev = n9
	n10.Next = nil
	n10.Child = nil

	n11.Val = 11
	n11.Prev = nil
	n11.Next = n12
	n11.Child = nil

	n12.Val = 12
	n12.Prev = n11
	n12.Next = nil
	n12.Child = nil

	root := flatten(n1)
	//var last *Node
	for root != nil {
		fmt.Printf("%v, %+v\n", &root.Next, root)
		if root.Next == nil {
			//last = root
		}
		root = root.Next
	}

	//for last != nil {
	//	fmt.Printf("%v, %+v\n", &last.Val, last)
	//	last = last.Prev
	//}

}

func TestRemoveNthFromEnd(t *testing.T) {
	var (
		t1 = new(ListNode)
		t2 = new(ListNode)
		t3 = new(ListNode)
		t4 = new(ListNode)
		t5 = new(ListNode)
	)
	t1.Val = 1
	t1.Next = t2
	t2.Val = 2
	t2.Next = t3
	t3.Val = 3
	t3.Next = t4
	t4.Val = 4
	t4.Next = t5
	t5.Val = 5
	node := removeNthFromEnd(t1, 2)
	for node != nil {
		fmt.Println(node.Val)
		node = node.Next
	}
}

func TestCalculate(t *testing.T) {
	fmt.Println(calculate("3/2 "))
}
