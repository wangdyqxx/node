/**
 * @Author: 李大双
 * @Description: 剑指 Offer 题解
 * @File: code_test
 * @Date: 2021/6/7 上午11:36
 */
package main

import "testing"

func TestDuplicate(t *testing.T) {
	testCase := []struct {
		Input  []int
		Output int
	}{
		{
			Input:  []int{2, 3, 1, 0, 2, 5},
			Output: 2,
		},
		{
			Input:  []int{2, 1, 3, 1, 4},
			Output: 1,
		},
	}

	for _, v := range testCase {
		if output := Duplicate(v.Input); output != v.Output {
			t.Errorf("expect %v,but is %v", v.Output, output)
		}
	}
}

func TestFind(t *testing.T) {
	matrix := [][]int{
		{1, 4, 7, 11, 15},
		{2, 5, 8, 12, 19},
		{3, 6, 9, 16, 22},
		{10, 13, 14, 17, 24},
		{18, 21, 23, 26, 30},
	}
	testCase := []struct {
		Target int
		Matrix [][]int
		Ret    bool
	}{
		{
			Target: 5,
			Matrix: matrix,
			Ret:    true,
		},
		{
			Target: 31,
			Matrix: matrix,
			Ret:    false,
		},
		{
			Target: 31,
			Matrix: [][]int{},
			Ret:    false,
		},
		{
			Target: 7,
			Matrix: [][]int{
				{1, 2, 8, 9},
				{4, 7, 10, 13},
			},
			Ret: true,
		},
	}

	for _, v := range testCase {
		if ret := Find(v.Target, v.Matrix); ret != v.Ret {
			t.Errorf("expect %v,but is %v", v.Ret, ret)
		}
	}
}

func TestReplaceSpace(t *testing.T) {
	testCase := []struct {
		Input  string
		Output string
	}{
		{
			Input:  "A B",
			Output: "A%20B",
		},
	}

	for _, v := range testCase {
		if output := ReplaceSpace(v.Input); output != v.Output {
			t.Errorf("expect %v,but is %v", v.Output, output)
		}
	}
}
