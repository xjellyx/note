package main

import "fmt"

func main() {
	var (
		dataArr = []int{3, 9, 7, 5, 61, 85, 99, 456, 21, 456}
	)
	fmt.Println(sortMaoBao(dataArr, true))
	fmt.Println(sortChoice(dataArr, true))
	fmt.Println(sortInsert(dataArr, true))
}

/*
	冒泡排序：对一个列表多次重复遍历。 比较相邻的两项，并且交换大小顺序排错的项。每对
列表实行一次遍历，就有一个最大项或者最小排在了正确的位置。总的来说，列表的每一个数据项都会在
其相应的位置“冒泡”。
	选择排序：比冒泡排序性能提高一点。它每遍历一次列表只交换一次数据，即进行一次遍历时找
到最大或者最小的项，完成遍历后，再把它换到正确的位置。
	插入排序：总是保持一个位置靠前的已排好的子表，然后每一个新的数据项被“插入”到前边的子表里，
排好的子表增加一项。
*/

// sortMaoBao sortCond true 大到小排序, 反之小到大
func sortMaoBao(data []int, sortCond bool) []int {
	var (
		l = len(data) - 1
	)
	for l > 0 {
		for i := 0; i < l; i++ {
			// 每一个当前位置与下一位进行比较，所有的数据都进行比较完毕，就可以得到排序后的数据
			if sortCond {
				if data[i+1] > data[i] {
					data[i+1], data[i] = data[i], data[i+1]

				}
			} else {
				if data[i+1] < data[i] {
					data[i+1], data[i] = data[i], data[i+1]

				}
			}
		}
		l--
	}
	return data
}

// sortChoice sortCond true 大到小排序, 反之小到大
func sortChoice(data []int, sortCond bool) []int {
	var (
		l = len(data) - 1
	)
	for l > 0 {
		var initIndex int
		// 初始值与第1个比较
		for i := 1; i < l+1; i++ {
			if sortCond {
				if data[i] < data[initIndex] {
					initIndex = i
				}
			} else {
				if data[i] > data[initIndex] {
					initIndex = i
				}
			}
		}
		data[l], data[initIndex] = data[initIndex], data[l]
		l--
	}
	return data
}

// sortInsert sortCond true 大到小排序, 反之小到大
func sortInsert(data []int, sortCond bool) []int {

	for l := 1; l < len(data); l++ {
		currentValue := data[l]
		position := l
		if !sortCond {
			for position > 0 && data[position-1] > currentValue {
				data[position] = data[position-1]
				position--
				data[position] = currentValue
			}
		} else {
			for position > 0 && data[position-1] < currentValue {
				data[position] = data[position-1]
				position--
				data[position] = currentValue
			}
		}
	}
	return data
}
