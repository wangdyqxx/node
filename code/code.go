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
