/**
 * @Author: 李大双
 * @Description: 剑指 Offer 题解
 * @File: array
 * @Date: 2021/6/7 上午11:18
 */
package code

import (
	"fmt"
	"math"
	"sort"
	"strconv"
	"strings"
)

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
				fmt.Println("d:", i, nums)
				return nums[i]
			}
			fmt.Println("a:", i, nums)
			nums[i], nums[nums[i]] = nums[nums[i]], nums[i]
			fmt.Println("b:", i, nums)
		}
		fmt.Println("c:", i, nums)
		nums[i], nums[nums[i]] = nums[nums[i]], nums[i]
	}
	return -1
}

func findRepeatNumber(arr []int) int {
	for i := range arr {
		for i != arr[i] {
			if arr[i] == arr[arr[i]] {
				fmt.Println("d:", i, arr)
				return arr[i]
			}
			fmt.Println("a:", i, arr)
			arr[i], arr[arr[i]] = arr[arr[i]], arr[i]
			fmt.Println("b:", i, arr)
		}
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

func findNumberIn2DArray(target int, matrix [][]int) bool {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return false
	}
	rows := len(matrix)
	for i := 0; i < rows; i++ {
		inums := matrix[i]
		ilen := len(inums)
		if inums[ilen-1] == target {
			return true
		}
		if inums[ilen-1] < target {
			continue
		}
		if inums[0] <= target && target < inums[ilen-1] {
			for _, inum := range inums {
				if inum == target {
					return true
				}
			}
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

func replaceSpace(s string) string {
	arr := []byte(s)
	tmp := make([]byte, len(arr)*3)
	i0 := byte(' ')
	i1 := byte('%')
	i2 := byte('2')
	i3 := byte('0')
	j := 0
	for i := 0; i < len(arr); i++ {
		v := arr[i]
		tmp[j] = v
		if v == i0 {
			tmp[j] = i1
			tmp[j+1] = i2
			tmp[j+2] = i3
			j += 3
		} else {
			j++
		}
	}
	return string(tmp[:j])
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

func spiralOrder(matrix [][]int) []int {
	var arr []int
	if len(matrix) == 0 {
		return arr
	}
	rows := len(matrix)
	cols := len(matrix[0])
	arr = make([]int, 0, rows*cols)
	r1, r2 := 0, rows-1
	c1, c2 := 0, cols-1
	for r1 <= r2 && c1 <= c2 {
		//上
		for i := c1; i <= c2; i++ {
			arr = append(arr, matrix[r1][i])
		}
		r1++
		//右
		for i := r1; i <= r2; i++ {
			arr = append(arr, matrix[i][c2])
		}
		c2--
		//下
		if r1 <= r2 {
			for i := c2; i >= c1; i-- {
				arr = append(arr, matrix[r2][i])
			}
			r2--
		}
		//左
		if c1 <= c2 {
			for i := r2; i >= r1; i-- {
				arr = append(arr, matrix[i][c1])
			}
			c1++
		}
	}
	return arr
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

func firstUniqChar(s string) byte {
	tmp := [26]int{}
	a := 'a'
	for _, b := range s {
		tmp[b-a]++
	}
	for _, b := range s {
		if tmp[b-a] == 1 {
			return byte(b)
		}
	}
	return ' '
}

//递增的三元子序列
func increasingTriplet(nums []int) bool {
	if len(nums) < 3 {
		return false
	}
	a := math.MaxInt32
	b := math.MaxInt32
	c := 0
	for i := 0; i < len(nums); i++ {
		c = nums[i]
		if a >= c {
			a = c
			continue
		}
		if b >= c {
			b = c
		}
		return true
	}
	return false
}

// 除自身以外数组的乘积
func productExceptSelf(nums []int) []int {
	size := len(nums)
	arr := make([]int, size)
	k := 1
	for i := 0; i < size; i++ {
		arr[i] = k
		k = k * nums[i]
	}
	k = 1
	for i := size - 1; i >= 0; i-- {
		arr[i] = arr[i] * k
		k = k * nums[i]
	}
	return arr
}

//搜索二维矩阵
func searchMatrix2(matrix [][]int, target int) bool {
	rows := len(matrix)
	cols := len(matrix[0])
	min := matrix[0][0]
	max := matrix[rows-1][cols-1]
	if target < min || max < target {
		return false
	}
	for i := 0; i < rows; i++ {
		if matrix[i][0] == target || target == matrix[i][cols-1] {
			return true
		}
		if matrix[i][0] < target && target < matrix[i][cols-1] && binarySearch(matrix[i], target) {
			return true
		}
	}
	return false
}

func binarySearch(nums []int, target int) bool {
	var a, b, c = 0, 0, len(nums) - 1
	for a <= c {
		b = (c-a)/2 + a
		if target == nums[b] {
			return true
		}
		if target > nums[b] {
			a = b + 1
		}
		if target < nums[b] {
			c = b - 1
		}
	}
	return false
}

//只出现一次的数字
func singleNumber(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}
	var k int
	for i := 0; i < len(nums); i++ {
		k = k ^ nums[i]
	}
	return k
}

//多数元素
func majorityElement(nums []int) int {
	mm := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		if _, ok := mm[nums[i]]; ok {
			mm[nums[i]]++
		} else {
			mm[nums[i]] = 1
		}
	}
	var num, value int
	for i, v := range mm {
		if v > value {
			num = i
			value = v
		}
	}
	return num
}

//搜索二维矩阵2
func searchMatrix(matrix [][]int, target int) bool {
	var cur int
	for x, y := len(matrix)-1, 0; x >= 0 && y < len(matrix[0]); {
		fmt.Println(x, y)
		cur = matrix[x][y]
		if cur == target {
			return true
		}
		if cur > target {
			x--
		} else {
			y++
		}
	}
	return false
}

//合并两个有序数组
func merge(nums1 []int, m int, nums2 []int, n int) {
	var k = len(nums1) - 1
	i, j := m-1, n-1
	for k >= 0 && (i >= 0 || j >= 0) {
		fmt.Println(k, i, j, ",", nums1)
		if i < 0 {
			nums1[k] = nums2[j]
			k--
			j--
			continue
		}
		if j < 0 {
			nums1[k] = nums1[i]
			k--
			i--
			continue
		}

		if nums2[j] >= nums1[i] {
			nums1[k] = nums2[j]
			k--
			j--
			continue
		}

		if nums2[j] < nums1[i] {
			nums1[k] = nums1[i]
			k--
			i--
		}

	}
}

//最大数
type Str []string

func (s Str) Len() int {
	return len(s)
}

func (s Str) Less(i, j int) bool {
	return fmt.Sprint(s[i], s[j]) > fmt.Sprint(s[j], s[i])
}

func (s Str) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
	return
}

func largestNumber(nums []int) string {
	arr := make([]string, 0, len(nums))
	for i := 0; i < len(nums); i++ {
		arr = append(arr, strconv.Itoa(nums[i]))
	}
	sort.Sort(Str(arr))
	var str strings.Builder
	if arr[0][0] == '0' {
		return "0"
	}
	for _, s := range arr {
		str.WriteString(s)
	}
	return str.String()
}

func lengthOfLongestSubstring(s string) int {
	var ans int
	bs := []byte(s)
	if len(bs) == 1 {
		return 1
	}
	mm := make(map[byte]struct{})
	for i, j := 0, 0; i < len(bs) && j <= i; {
		if _, ok := mm[bs[i]]; ok {
			newSize := len(mm)
			if newSize > ans {
				ans = newSize
			}
			delete(mm, bs[j])
			j++
		} else {
			mm[bs[i]] = struct{}{}
			i++
		}
	}
	newSize := len(mm)
	if newSize > ans {
		ans = newSize
	}
	return ans
}

//盛水最多的容器
func maxArea(height []int) int {
	size := len(height)
	var area, tmp int
	var y int
	for i, j := 0, size-1; i < j; {
		com := height[i] >= height[j]
		if com {
			y = height[j]
			tmp = y * (j - i)
			j--
		} else {
			y = height[i]
			tmp = y * (j - i)
			i++
		}
		if tmp > area {
			area = tmp
		}
	}
	return area
}

//三数之和
func threeSum(nums []int) [][]int {
	ans := make([][]int, 0)
	sort.Ints(nums)
	size := len(nums)
	if size < 3 {
		return ans
	}
	//if nums[0] > 0 {
	//	return ans
	//}
	fmt.Println(nums)
	for a := 0; a < size; a++ {
		if a > 0 && nums[a] == nums[a-1] {
			continue
		}
		target := -1 * nums[a]
		c := size - 1
		for b := a + 1; b < size; b++ {
			if b > a+1 && nums[b] == nums[b-1] {
				continue
			}
			//fmt.Println(a, b, c)
			if target > nums[b]+nums[c] {
				continue
			}
			for b < c && target < nums[b]+nums[c] {
				c--
			}
			if b == c {
				break
			}
			if nums[a]+nums[b]+nums[c] == 0 {
				ans = append(ans, []int{nums[a], nums[b], nums[c]})
			}
		}
	}
	return ans
}

//零钱兑换
func coinChange(coins []int, amount int) int {
	memo := make([]int, amount+1)
	for i := 0; i < len(memo); i++ {
		memo[i] = math.MaxInt32
	}
	ans := recCoinChange(coins, memo, amount)
	fmt.Println(memo)
	return ans
}

func recCoinChange(coins []int, memo []int, amount int) int {
	if amount == 0 {
		return 0
	}
	if amount < 0 {
		return -1
	}
	if memo[amount] != math.MaxInt32 {
		return memo[amount]
	}
	ans := math.MaxInt32
	for i := 0; i < len(coins); i++ {
		if amount-coins[i] < 0 {
			continue
		}
		tmp := recCoinChange(coins, memo, amount-coins[i])
		if tmp == -1 {
			continue
		}
		ans = min(ans, tmp+1)
	}
	if ans == math.MaxInt32 {
		ans = -1
	} else {
		memo[amount] = ans
	}
	return ans
}

func Dp_CoinChange(coins []int, amount int) int {
	dp := make([]int, amount+1)
	dp[0] = 0
	for i := 1; i <= amount; i++ {
		dp[i] = math.MaxInt32
		for j := 0; j < len(coins); j++ {
			if coins[j] <= i {
				dp[i] = min(dp[i], dp[i-coins[j]]+1)
			}
		}
	}
	if dp[amount] == math.MaxInt32 {
		return -1
	}
	return dp[amount]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}

func fib1(n int) int {
	arr := make([]int, n)
	arr[0] = 1
	arr[1] = 1
	num := 2
	for num < n {
		arr[num] = arr[num-1] + arr[num-2]
		num++
	}
	return arr[len(arr)-1]
}

func fib2(n int) int {
	var (
		a = 1
		b = 1
		c = 0
	)
	for n >= 3 {
		n--
		c = a + b
		a = b
		b = c
	}
	return c
}

func fib3(n int) int {
	if n <= 2 {
		return 1
	}
	return fib3(n-1) + fib3(n-2)
}

//打家劫舍
func rob(nums []int) int {
	//return recRob(nums, len(nums)-1)
	return dpRob(nums)
}

func recRob(nums []int, index int) int {
	if index == 1 || index == 0 {
		return max(nums[index], nums[0])
	}
	return max(recRob(nums, index-2)+nums[index], recRob(nums, index-1))
}

func dpRob(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}
	dp := make([]int, len(nums))
	dp[0] = nums[0]
	dp[1] = max(nums[0], nums[1])
	for i := 2; i < len(nums); i++ {
		dp[i] = +max(nums[i]+dp[i-2], dp[i-1])
	}
	return dp[len(nums)-1]
}

func lengthOfLIS(nums []int) int {
	return dpLengthOfLIS(nums)
}

func dpLengthOfLIS(nums []int) int {
	arr := make([]int, len(nums))
	for i := 0; i < len(nums); i++ {
		arr[i] = 1
		for j := 0; j < i; j++ {
			if nums[j] < nums[i] {
				arr[i] = max(arr[i], arr[j]+1)
			}
		}
	}
	fmt.Println(arr)
	var ans int
	for i := 0; i < len(arr); i++ {
		ans = max(ans, arr[i])
	}
	return ans
}

func arr1(arr []int) {
	var tmp []int
	arr[0] = 2
	copy(tmp, arr)
	fmt.Println("a1:", arr)
	return
}

func arr2(arr []int) {
	arr[0] = 2
	fmt.Println("a2:", arr)
	return
}
