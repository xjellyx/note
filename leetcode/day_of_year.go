package main

import (
	"fmt"
	"strconv"
	"strings"
)

var (
	m = map[int]int{
		1: 31,
		3: 31, 4: 30, 5: 31, 6: 30, 7: 31, 8: 31, 9: 30, 10: 31, 11: 30, 12: 30,
	}
)

const (
	runTwoMonth = 29
	twoMonth    = 28
)

func dayOfYear(date string) (res int) {
	var isRunYear = false
	arrStr := strings.Split(date, "-")
	year, _ := strconv.Atoi(arrStr[0])

	if year%4 == 0 {
		if year%100 == 0 && year%400 != 0 {
			isRunYear = false
		}
		isRunYear = true
	}
	month, _ := strconv.Atoi(arrStr[1])
	for i := 1; i < month; i++ {
		if i != 2 {
			res += m[i]
		} else {
			if isRunYear {
				res += runTwoMonth
			} else {
				res += twoMonth
			}
		}
	}
	day, _ := strconv.Atoi(arrStr[2])
	res += day
	return
}

func main() {
	fmt.Println(dayOfYear("2004-03-01"))
}
