package main

import (
	"github.com/olongfen/note/data"
)

func main() {
	o, _ := data.Open(data.Config{Mode: "http"})
	o.GetData()
}
