package main

import (
	"fmt"
	"os/exec"
)

func main() {
	c := exec.Command("go", "run", "/data/gocode/src/github.com/olongfen/note/main/fucn_sort_learn.go", "&&", "go", "env")
	fmt.Println(c.String())
	out, err := c.Output()
	fmt.Println(string(out), err)
}
