package main

import (
	"container/list"
	"fmt"
	"math"
	"sort"
	"strconv"
	"sync"
)

// sortedSquares 有序数组的平方
func sortedSquares(nums []int) []int {
	var (
		newNums = make([]int, len(nums))
	)
	for i, v := range nums {
		newNums[i] = v * v
	}

	sort.Ints(newNums)
	return newNums
}

func rotate(nums []int, k int) {
	newNums := make([]int, len(nums))
	for i, v := range nums {
		newNums[(i+k)%len(nums)] = v
	}
	copy(nums, newNums)
	fmt.Println(nums)
}

func moveZeroes(nums []int) {
	l := len(nums)
	left, right := 0, 0
	for right < l {
		if nums[right] != 0 {
			nums[left], nums[right] = nums[right], nums[left]
			left++
		}
		right++
	}

}

func twoSum(numbers []int, target int) []int {
	l := len(numbers)
	var res []int
	for i := 0; i < l-1; i++ {
		for j := i + 1; j < l; j++ {
			if numbers[i]+numbers[j] == target {
				res = append(res, i+1, j+1)
			}
		}
	}
	return res
}

func judgeSquareSum(c int) bool {
	var (
		array = make([]int, 0)
	)
	for i := 0; i*i <= c; i++ {
		array = append(array, i)
	}
	var (
		l    = len(array)
		slow = l - 1
		fast = 0
	)
	for fast <= slow {
		v1 := array[fast] * array[fast]
		v2 := array[slow] * array[slow]
		if v1+v2 == c {
			return true
		} else if v1+v2 > c {
			slow--
		} else {
			fast++
		}

	}
	return false
}

func totalFruit(fruits []int) int {
	var (
		l     = len(fruits)
		slow  = 0
		fast  = 0
		res   = 0
		count = 0
		m     = make(map[int]int, 0)
	)
	for fast < l && slow <= fast {
		fmt.Println(slow, fast, m, count)
		if _, ok := m[fruits[fast]]; ok {
			m[fruits[fast]] += 1
			count++
			fast++
			if count > res {
				res = count
			}
		} else if len(m) < 2 {
			m[fruits[fast]] += 1
			count++
			fast++
			if count > res {
				res = count
			}
		} else {
			slow++
			fast = slow
			count = 0
			m = make(map[int]int, 0)
		}
	}

	return res
}

func getNext(next []int, s string) {
	var (
		j = 0 // 后缀末尾位置
	)
	next[0] = 0
	for i := 1; i < len(s); i++ { // 前缀末尾位置
		for j > 0 && s[i] != s[j] {
			j = next[j-1]
		}
		if s[i] == s[j] {
			j++
		}
		next[i] = j
	}
}

func repeatedSubstringPattern(s string) bool {
	var (
		l    = len(s)
		next = make([]int, l)
	)

	if l == 0 {
		return false
	}

	getNext(next, s)
	if next[l-1] != 0 && l%(l-next[l-1]) == 0 {
		return true
	}
	return false
}

func strStr(haystack string, needle string) int {
	var (
		l    = len(needle)
		j    = 0
		next = make([]int, l)
	)
	getNext(next, needle)
	for i := 0; i < len(haystack); i++ {
		for j > 0 && needle[j] != haystack[i] {
			j = next[j-1]
		}
		if needle[j] == haystack[i] {
			j++
		}
		if j == l {
			return i - l + 1
		}
	}

	return -1
}

func generateMatrix(n int) [][]int {
	top, bottom := 0, n-1
	left, right := 0, n-1
	num := 1
	tar := n * n
	matrix := make([][]int, n)
	for i := 0; i < n; i++ {
		matrix[i] = make([]int, n)
	}
	for num <= tar {
		for i := left; i <= right; i++ {
			matrix[top][i] = num
			num++
		}
		top++
		for i := top; i <= bottom; i++ {
			matrix[i][right] = num
			num++
		}
		right--
		for i := right; i >= left; i-- {
			matrix[bottom][i] = num
			num++
		}
		bottom--
		for i := bottom; i >= top; i-- {
			matrix[i][left] = num
			num++
		}
		left++
	}
	return matrix
}

func singleNumber(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}
	for i := 2; i < len(nums); {
		if nums[i] != nums[i-1] {
			return nums[i]
		} else {
			i += 2
		}
	}

	return 0
}

func evalRPN(tokens []string) int {
	var (
		l = list.New()
	)
	for _, v := range tokens {

		switch v {
		case "+":
			v2 := l.Front()
			l.Remove(v2)
			v3 := l.Front()
			l.Remove(v3)
			l.PushFront(v2.Value.(int) + v3.Value.(int))

		case "-":
			v2 := l.Front()
			l.Remove(v2)
			v3 := l.Front()
			l.Remove(v3)
			l.PushFront(v3.Value.(int) - v2.Value.(int))
		case "*":
			v2 := l.Front()
			l.Remove(v2)
			v3 := l.Front()
			l.Remove(v3)
			l.PushFront(v2.Value.(int) * v3.Value.(int))
		case "/":
			v2 := l.Front()
			l.Remove(v2)
			v3 := l.Front()
			l.Remove(v3)
			l.PushFront(v3.Value.(int) / v2.Value.(int))
		default:
			_v, _ := strconv.Atoi(v)
			l.PushFront(_v)
		}
	}
	return l.Front().Value.(int)
}

func maxSlidingWindow(nums []int, k int) []int {
	var (
		i   = 0
		j   = k
		arr []int
		wg  = sync.WaitGroup{}
	)
	for j <= len(nums) {
		wg.Add(1)
		go func() {
			defer wg.Done()
			l := j - i
			dn := make([]int, l)
			copy(dn, nums[i:j])
			sort.Ints(dn)
			arr = append(arr, dn[len(dn)-1])
			i++
			j++
		}()
	}
	wg.Wait()
	return arr
}

func chalkReplacer(chalk []int, k int) int {
	l := len(chalk) - 1
	_l := 0
	index := -1
	for {
		if _l > l {
			_l = 0
		}
		k -= chalk[_l]
		fmt.Println(k, _l)
		if k < 0 {
			index = _l
			break
		}

		_l++
	}
	return index
}

func main() {
	fmt.Println(math.Abs(-1))
	fmt.Println(1 << 4)
	fmt.Println(bitCount(2 ^ 4))
	fmt.Println(totalHammingDistance([]int{4, 14, 4}))
}

func totalHammingDistance(nums []int) int {
	sum := 0
	left := 0
	l := len(nums)
	right := 1
	for left < right && left < l-1 {
		sum = sum + bitCount(nums[left]^nums[right])
		right++
		if right == l {
			left++
			right = left + 1

		}

	}
	return sum
}

func bitCount(n int) int {
	var (
		c = 0
	)
	for n != 0 {
		c++
		n &= n - 1
	}
	return c
}
