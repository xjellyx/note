package main

import (
	"fmt"
	"sort"
	"strings"
)

func isAnagram(s string, t string) bool {
	var (
		m = make(map[byte]int)
	)
	for i := 0; i < len(s); i++ {
		if val, ok := m[s[i]]; !ok {
			m[s[i]] = 1
		} else {
			m[s[i]] = val + 1
		}
	}

	for i := 0; i < len(t); i++ {
		if val, ok := m[t[i]]; !ok {
			return false
		} else {
			m[t[i]] = val - 1
		}
	}

	for _, v := range m {
		if v > 0 {
			return false
		}
	}
	return true
}

func intersection(nums1 []int, nums2 []int) []int {
	var (
		m   = make(map[int]bool)
		res []int
	)
	for _, v := range nums1 {
		m[v] = true
	}
	for _, v := range nums2 {
		if _, ok := m[v]; ok {
			res = append(res, v)
			delete(m, v)
		}
	}
	return res
}

func intersect(nums1 []int, nums2 []int) []int {
	var (
		m   = make(map[int]int)
		res []int
	)
	for _, v := range nums1 {
		if val, ok := m[v]; ok {
			m[v] = val + 1
		} else {
			m[v] = 1
		}
	}
	for _, v := range nums2 {
		if val, ok := m[v]; ok && val > 0 {
			res = append(res, v)
			m[v] = val - 1
		}
	}
	return res
}

func isHappy(n int) bool {
	sum := n
	var (
		m = make(map[int]bool)
	)
	for sum != 1 {
		if _, ok := m[sum]; !ok {
			m[sum] = true
		} else {
			return false
		}
		sum = getSum(sum)
	}
	return sum == 1
}

func getSum(n int) int {
	sum := 0
	for n > 0 {
		sum += (n % 10) * (n % 10)
		n = n / 10
	}
	return sum
}

func twoSum3(nums []int, target int) []int {
	var (
		m   = make(map[int]int)
		res []int
	)
	for i, v := range nums {
		if val, ok := m[v]; !ok {
			m[target-v] = i
		} else {
			res = append(res, val, i)
			return res
		}
	}
	return res
}

func canConstruct(ransomNote string, magazine string) bool {
	var (
		arr = make([]int, 26)
	)
	for i := 0; i < len(magazine); i++ {
		arr[magazine[i]-'a'] += 1
	}
	for i := 0; i < len(ransomNote); i++ {
		arr[ransomNote[i]-'a'] -= 1

		if arr[ransomNote[i]-'a'] < 0 {
			return false
		}
	}
	return true
}

func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	var (
		l     = len(nums)
		right = l - 1
		left  = 0
		res   [][]int
	)
	for i := 0; i < l-2; i++ {
		if nums[i] > 0 {
			return res
		}

		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		left = i + 1
		right = l - 1
		for left < right {
			n2, n3 := nums[left], nums[right]
			sum := nums[i] + n2 + n3
			if sum > 0 {
				right--
			} else if sum < 0 {
				left++
			} else {
				d := []int{nums[i], n2, n3}
				res = append(res, d)
				for left < right && nums[left] == n2 {
					left++
				}
				for left < right && nums[right] == n3 {
					right--
				}
			}
		}

	}

	return res
}

func main() {
	fmt.Println(threeSum([]int{0, 0, 0, 0}))
}
