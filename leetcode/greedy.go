package main

import (
	"fmt"
	"sort"
)

func findContentChildren(g []int, s []int) int {
	sort.Ints(g)
	sort.Ints(s)
	index := len(s) - 1
	var res int
	for i := len(g) - 1; i >= 0; i-- {
		if index >= 0 && s[index] >= g[i] {
			res++
			index--
		}
	}
	return res
}

func numRescueBoats(people []int, limit int) int {
	sort.Ints(people)
	l := len(people)
	var (
		res  int
		last int
	)

	for i := l - 1; i > 0; i-- {
		if people[i] >= limit {
			res++
		} else {
			if last > 0 {
				if last-people[i] == 0 {
					res++
				} else if last-people[i] > 0 {
					last = last - people[i]
				} else {
					res++
				}
			} else {
				last = limit - people[i]
			}

		}
	}
	return res
}

func main() {
	fmt.Println(numRescueBoats([]int{1, 2}, 3))
}
