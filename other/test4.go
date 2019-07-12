package main

import (
	"fmt"
	"strings"
)

func main() {
	arr := []string{"sdsd", "asrgfd", "fjijokitj"}
	s := strings.Join(arr, "\n")
	fmt.Println(s)
}
