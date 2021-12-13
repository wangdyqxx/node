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

type NestedInteger struct {
}

func (this NestedInteger) IsInteger() bool {
	return false
}
func (this NestedInteger) GetInteger() int {
	return 0
}
func (this *NestedInteger) SetInteger(value int) {}
func (this *NestedInteger) Add(elem NestedInteger) {

}

func (this NestedInteger) GetList() []*NestedInteger {
	return nil
}

func Constructor1(nestedList []*NestedInteger) *NestedIterator {
	iter := new(NestedIterator)
	var dfs func(list []*NestedInteger)
	dfs = func(list []*NestedInteger) {
		for _, integer := range list {
			if integer.IsInteger() {
				iter.values = append(iter.values, integer.GetInteger())
			} else {
				dfs(integer.GetList())
			}
		}
	}
	dfs(nestedList)
	iter.size = len(iter.values)
	return iter
}

type NestedIterator struct {
	index  int
	size   int
	values []int
}

func (this *NestedIterator) Next() int {
	num := this.values[this.index]
	this.index++
	return num
}

func (this *NestedIterator) HasNext() bool {
	if this.index >= len(this.values) {
		return false
	}
	return true
}
