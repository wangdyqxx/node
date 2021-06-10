/**
 * @Author: 李大双
 * @Description:
 * @File: stack
 * @Date: 2021/6/10 上午9:47
 */
package code

var stack1 []int
var stack2 []int

func Push(node int) {
	stack1 = append(stack1, node)
}

func Pop() int {
	if len(stack2) == 0 {
		l := len(stack1)
		for i := l - 1; i >= 0; i-- {
			stack2 = append(stack2, stack1[i])
		}
		stack1 = []int{}
	}
	if len(stack2) == 0 {
		return -1
	}
	node := stack2[len(stack2)-1]
	stack2 = stack2[:len(stack2)-1]
	return node
}
