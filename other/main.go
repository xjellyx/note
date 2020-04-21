package main

import (
	"fmt"
	"strconv"
	"sync"
	"time"
)

type data struct {
	Msg map[string]interface{}
}

func main() {
	p := &sync.Pool{}
	p.New = func() interface{} {
		return &data{}
	}

	for i := 0; i < 10; i++ {
		go func(_i int) {
			da := &data{
				Msg: map[string]interface{}{
					"index": strconv.Itoa(_i),
				},
			}
			p.Put(da)
		}(i)
	}
	time.Sleep(time.Millisecond * 29)
	var (
		i int
	)
	for i < 5 {
		i++
		fmt.Println(p.Get(), i)
	}

}
