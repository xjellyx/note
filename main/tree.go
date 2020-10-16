package main

import (
	"fmt"
	"strings"
	"time"
)

func main() {
	fmt.Println(strings.HasPrefix("dr:23123", "dr:"))
	fmt.Println(strings.Split("dr:23123", "dr:"))
	fmt.Println(time.Parse("2006", "2018"))
	fmt.Println("%" + fmt.Sprintf(`%s`, "address") + "%")
}
