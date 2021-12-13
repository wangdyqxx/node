/*
 @Author : dongyang
 @Time   : 2021/9/17 17
 @Desc   :
*/
package main

import (
	"fmt"
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func main() {
	var arr []int
	for i := 10; i < 20; i++ {
		arr = append(arr, i)
	}
	fmt.Println(arr)
	fmt.Println(search(arr, 18))
}

// 二分搜索最常用模板
func search(nums []int, target int) int {
	// 1、初始化start、end
	start := 0
	end := len(nums) - 1
	// 2、处理for循环
	for start+1 < end {
		mid := start + (end-start)/2
		// 3、比较a[mid]和target值
		if nums[mid] == target {
			end = mid
		} else if nums[mid] < target {
			start = mid
		} else if nums[mid] > target {
			end = mid
		}
	}
	// 4、最后剩下两个元素，手动判断
	if nums[start] == target {
		return start
	}
	if nums[end] == target {
		return end
	}
	return -1
}
