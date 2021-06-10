/**
 * @Author: 李大双
 * @Description:
 * @File: stack_test
 * @Date: 2021/6/10 上午9:48
 */
package code

import (
	"fmt"
	"testing"
)

func TestPop(t *testing.T) {
	stack1 = []int{}
	stack2 = []int{}
	Push(1)
	Push(2)
	Push(3)
	fmt.Println(Pop())
	fmt.Println(Pop())
}
