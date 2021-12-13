/*
 @Author : dongyang
 @Time   : 2021/9/24 16
 @Desc   :
*/
package code

import (
	"fmt"
	"strconv"
	"strings"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

//type TreeNode struct {
//	Val   int
//	Left  *TreeNode
//	Right *TreeNode
//}

//前序递归
func preorderTraversal1(tree *TreeNode) {
	if tree == nil {
		return
	}
	fmt.Println(tree.Val)
	preorderTraversal1(tree.Left)
	preorderTraversal1(tree.Right)
}

//前序非递归 根左右
func preorderTraversal2(root *TreeNode) []int {
	var arr []int
	if root == nil {
		return nil
	}
	var stack []*TreeNode
	for root != nil || len(stack) > 0 {
		// 前序遍历，所以先保存结果
		for root != nil {
			arr = append(arr, root.Val)
			stack = append(stack, root)
			root = root.Left
		}
		tmp := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		root = tmp.Right
	}
	return arr
}

//中序非递归 左根右
func inorderTraversal(root *TreeNode) []int {
	var arr []int
	if root == nil {
		return nil
	}

	stack := make([]*TreeNode, 0)
	for len(stack) > 0 || root != nil {
		for root != nil {
			stack = append(stack, root)
			//fmt.Println("input：", root.Val)
			root = root.Left // 一直向左
		}
		// 弹出
		val := stack[len(stack)-1]
		arr = append(arr, val.Val)
		stack = stack[:len(stack)-1]
		root = val.Right
	}
	return arr
}

//后序非递归 左右根
func postorderTraversal1(root *TreeNode) []int {
	var arr []int
	if root == nil {
		return arr
	}
	var tmp *TreeNode
	stack := make([]*TreeNode, 0)
	for len(stack) > 0 || root != nil {
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}
		// 弹出
		root = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if root.Right == nil || tmp == root.Right {
			arr = append(arr, root.Val)
			tmp = root
			root = nil
		} else {
			stack = append(stack, root)
			root = root.Right
		}
	}
	return arr
}

//后序非递归 根右左 翻转-> 左右根
func postorderTraversal2(root *TreeNode) []int {
	var arr, tmp []int
	if root == nil {
		return arr
	}
	stack := make([]*TreeNode, 0)
	for len(stack) > 0 || root != nil {
		for root != nil {
			stack = append(stack, root)
			tmp = append(tmp, root.Val)
			root = root.Right // 一直向左
		}
		// 弹出
		val := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		root = val.Left
	}
	for i := len(tmp) - 1; i >= 0; i-- {
		arr = append(arr, tmp[i])
	}
	return arr
}

// V1：深度遍历，结果指针作为参数传入到函数内部
func postorderTraversal3(root *TreeNode) []int {
	result := make([]int, 0)
	dfs(root, &result)
	return result
}

func dfs(root *TreeNode, result *[]int) {
	if root == nil {
		return
	}
	*result = append(*result, root.Val)
	dfs(root.Left, result)
	dfs(root.Right, result)
}

func postorderTraversal4(root *TreeNode) []int {
	var arr []int
	if root == nil {
		return arr
	}
	var stack []*TreeNode
	for root != nil || len(stack) != 0 {
		for root != nil {
			stack = append(stack, root)
			arr = append(arr, root.Val)
			root = root.Right
		}
		root = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		root = root.Left
	}
	for i, j := 0, len(arr)-1; i < j; i++ {
		arr[i], arr[j] = arr[j], arr[i]
		j--
	}
	return arr
}

//最大深度
func maxDepth(root *TreeNode) int {
	if root != nil {
		return 0
	}
	var depth1, depth2 int
	depth(root, depth1, depth2)
	if depth1 > depth2 {
		return depth1
	}
	if depth1 < depth2 {
		return depth2
	}
	return depth2
}

//    3
//   / \
//  9  20
//    /  \
//   15   7
func depth(root *TreeNode, depth1, depth2 int) {
	fmt.Println(depth1, depth2)
	if root.Left != nil {
		depth1++
		depth(root.Left, depth1, depth2)
	}
	if root.Right != nil {
		depth2++
		depth(root.Right, depth1, depth2)
	}
	return
}

//二叉搜索树中第K小的元素
func kthSmallest(root *TreeNode, k int) int {
	var stack []*TreeNode
	for root != nil || len(stack) > 0 {
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}
		tmp := stack[len(stack)-1]
		k--
		fmt.Println(tmp)
		if k == 0 {
			return tmp.Val
		}
		root = tmp.Right
		stack = stack[:len(stack)-1]
	}
	return 0
}

//二叉树的最近公共祖先
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == nil {
		return root
	}
	if root == p || root == q {
		return root
	}
	left := lowestCommonAncestor(root.Left, p, q)
	right := lowestCommonAncestor(root.Right, p, q)
	if left == nil && right == nil {
		return nil
	}
	if left == nil {
		return right
	}
	if right == nil {
		return left
	}
	return root
}

//二叉树的序列化与反序列化
type Codec struct {
}

func Constructor() Codec {
	return Codec{}
}

// Serializes a tree to a single string.
func (this *Codec) serialize(root *TreeNode) string {
	var bs strings.Builder
	var dfs func(root *TreeNode)
	dfs = func(root *TreeNode) {
		if root == nil {
			bs.WriteString("x,")
			return
		}
		bs.WriteString(strconv.Itoa(root.Val) + ",")
		dfs(root.Left)
		dfs(root.Right)
		return
	}
	dfs(root)
	return bs.String()
}

// Deserializes your encoded data to tree.
func (this *Codec) deserialize(data string) *TreeNode {
	strs := strings.Split(data, ",")
	var desc func() *TreeNode
	desc = func() *TreeNode {
		if len(strs) == 0 {
			return nil
		}
		str := strs[0]
		strs = strs[1:]
		if str == "x" {
			return nil
		}
		val, _ := strconv.Atoi(str)
		node := &TreeNode{
			Val:   val,
			Left:  desc(),
			Right: desc(),
		}
		return node
	}
	return desc()
}

//层序遍历二叉树
func levelOrder111(root *TreeNode) [][]int {
	ret := [][]int{}
	if root == nil {
		return ret
	}
	q := []*TreeNode{root}
	for i := 0; len(q) > 0; i++ {
		ret = append(ret, []int{})
		p := []*TreeNode{}
		for j := 0; j < len(q); j++ {
			node := q[j]
			ret[i] = append(ret[i], node.Val)
			if node.Left != nil {
				p = append(p, node.Left)
			}
			if node.Right != nil {
				p = append(p, node.Right)
			}
		}
		q = p
	}
	return ret
}

func levelOrder(root *TreeNode) [][]int {
	var arrres [][]int
	if root == nil {
		return arrres
	}
	ts := []*TreeNode{root}
	for i := 0; i < len(ts); i++ {
		var ts1 []*TreeNode
		arrres = append(arrres, []int{})
		for j := 0; j < len(ts); j++ {
			arrres[i] = append(arrres[i], ts[j].Val)
			if ts[j].Left != nil {
				ts1 = append(ts1, ts[j].Left)
			}
			if ts[j].Right != nil {
				ts1 = append(ts1, ts[j].Right)
			}
		}
		ts = ts1
	}
	return arrres
}
