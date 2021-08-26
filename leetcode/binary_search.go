package main

// search 二分查找
func search(nums []int, target int) int {
	var (
		right = len(nums) - 1
		left  = 0
		mid   int
	)
	for left <= right {
		mid = left + (right-left)>>1
		if nums[mid] == target {
			return mid
		} else if nums[mid] > target {
			right = mid - 1
		} else if nums[mid] < target {
			left = mid + 1
		}

	}
	return -1
}

// searchInsert 搜索插入位置
func searchInsert(nums []int, target int) int {
	var (
		left, right int
		mid         int
		index       int
	)
	right = len(nums) - 1
	for left <= right {
		mid = left + (right-left)>>1
		if nums[mid] > target {
			right = mid - 1
			if right < 0 {
				index = 0
				continue
			}
			if target > nums[right] {
				index = right + 1
				continue
			}
			if target < nums[right] && target > nums[left] {
				index = right
				continue
			}
			if target < nums[left] {
				index = left
				continue
			}
		} else if nums[mid] < target {
			left = mid + 1
			if target > nums[right] {
				index = right + 1
				continue
			}
			if target < nums[right] && target > nums[left] {
				index = right
				continue
			}

			if target < nums[left] {
				index = left
				continue
			}
		} else if nums[mid] == target {
			return mid
		}

	}
	return index

}
