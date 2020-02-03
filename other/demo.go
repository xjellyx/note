package main

import (
	"fmt"
	"github.com/chenhg5/collection"
	"sort"
)

type IntSlice []int

func (c IntSlice) Len() int {
	return len(c)
}

func (c IntSlice) Swap(i, j int) {
	c[i], c[j] = c[j], c[i]
}

func (c IntSlice) Less(i, j int) bool {
	return c[i] < c[j]
}

func main() {
	array := IntSlice{5, 5, 9, 0}
	a := collection.Collect([]int{1, 9, 1, 0, 9}).Sum()
	fmt.Println(a.IntPart())
	target := 10
	sort.Sort(array)
	var result [3]int

	for i := 0; i < len(array)-2; i++ {
		if i > 0 && array[i-1] == array[i] {
			continue
		}
		j := i + 1
		k := len(array) - 1

		for j < k {
			if array[i]+array[j]+array[k] > target {
				k--
				for array[k] == array[k+1] && j < k {
					k--
				}
			} else if array[i]+array[j]+array[k] < target {
				j++
				for array[j] == array[j-1] && j < k {
					j++
				}
			} else {
				result = [3]int{array[i], array[j], array[k]}
				j++
				k--
				for array[k] == array[k+1] && j < k {
					k--
				}
				for array[j] == array[j-1] && j < k {
					j++
				}
			}
		}

	}

	fmt.Println(result)

}
