/*
 @Author : dongyang
 @Time   : 2021/9/24 16
 @Desc   :
*/
package code

import (
	"fmt"
	"testing"
)
/**
		 1
	 2		  5
  3	    4	6	7
8   9

 */

func TestTree_preorderTraversal(t *testing.T) {
	t1 := &TreeNode{Val: 1}
	t8, t9 := &TreeNode{Val: 8}, &TreeNode{Val: 9}
	t3 := &TreeNode{
		Val:   3,
		Left:  t8,
		Right: t9,
	}
	t4 := &TreeNode{Val: 4}
	t2 := &TreeNode{Val: 2,
		Left:  t3,
		Right: t4,
	}
	t1.Left = t2
	t6 := &TreeNode{Val: 6}
	t7 := &TreeNode{Val: 7}
	t5 := &TreeNode{
		Val:   5,
		Left:  t6,
		Right: t7,
	}
	t1.Right = t5
	fmt.Println(preorderTraversal2(t1))
	//fmt.Println(inorderTraversal(t1))
	//postorderTraversal1(tree)
	//arr := postorderTraversal2(tree)
	//arr := postorderTraversal4(tree)
	//fmt.Println(arr)
	//val := kthSmallest(tree, 3)
	//val := lowestCommonAncestor(tree, t8, t7)
	//codec := Constructor()
	//val := codec.serialize(t1)
	//t.Log(val)
	//t1qq:=codec.deserialize(val)
	//t.Log(t1qq)
	//valqq := codec.serialize(t1qq)
	//t.Log(valqq)
}
