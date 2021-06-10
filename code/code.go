/**
 * @Author: 李大双
 * @Description: 剑指 Offer 题解
 * @File: code
 * @Date: 2021/6/7 上午11:18
 */
package main

/*
 * @Author 李大双
 * @Description: 数组中重复的数字
 * @param: nums
 * @return: int
 */
func Duplicate(nums []int) int {
	for i := 0; i < len(nums); i++ {
		for i != nums[i] {
			if nums[i] == nums[nums[i]] {
				return nums[i]
			}
			nums[i], nums[nums[i]] = nums[nums[i]], nums[i]
		}
		nums[i], nums[nums[i]] = nums[nums[i]], nums[i]
	}
	return -1
}

/*
 * @Author 李大双
 * @Description: 二维数组中的查找
 * @param: target
 * @param: matrix
 * @return: bool
 */
func Find(target int, matrix [][]int) bool {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return false
	}
	rows := len(matrix)
	cols := len(matrix[0])
	r := 0
	c := cols - 1
	for r <= rows-1 && c >= 0 {
		if target == matrix[r][c] {
			return true
		} else if target > matrix[r][c] {
			r++
		} else {
			c--
		}
	}
	return false
}

/*
 * @Author 李大双
 * @Description: 替换空格
 * @param: s
 * @return: string
 */
func ReplaceSpace(s string) string {
	array := []byte(s)
	p1 := len(array) - 1
	for _, v := range array {
		if v == ' ' {
			array = append(array, ' ', ' ')
		}
	}
	p2 := len(array) - 1
	for p1 >= 0 && p2 > p1 {
		c := array[p1]
		p1--
		if c == ' ' {
			array[p2] = '0'
			p2--
			array[p2] = '2'
			p2--
			array[p2] = '%'
			p2--
		} else {
			array[p2] = c
			p2--
		}
	}
	return string(array)
}

/*
 * @Author 李大双
 * @Description: 顺时针打印矩阵
 * @param: matrix
 * @return: []int
 */
func PrintMatrix(matrix [][]int) []int {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return nil
	}
	l := len(matrix) * len(matrix[0])
	ret := make([]int, 0, l)
	r1, r2, c1, c2 := 0, len(matrix)-1, 0, len(matrix[0])-1
	for r1 <= r2 && c1 <= c2 {
		// 上
		for i := c1; i <= c2; i++ {
			ret = append(ret, matrix[r1][i])
		}
		// 右
		for i := r1 + 1; i <= r2; i++ {
			ret = append(ret, matrix[i][c2])
		}
		// 下
		if r1 != r2 {
			for i := c2 - 1; i >= c1; i-- {
				ret = append(ret, matrix[r2][i])
			}
		}
		// 左
		if c1 != c2 {
			for i := r2 - 1; i > r1; i-- {
				ret = append(ret, matrix[i][c1])
			}
		}
		r1++
		r2--
		c1++
		c2--
	}
	return ret
}

/*
 * @Author 李大双
 * @Description: 第一个只出现一次的字符位置
 * @param: str
 * @return: int
 */
func FirstNotRepeatingChar(str string) int {
	array := make([]int, 128)
	for _, r := range str {
		array[r]++
	}
	for i, r := range str {
		if array[r] == 1 {
			return i
		}
	}
	return -1
}

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
