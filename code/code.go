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
