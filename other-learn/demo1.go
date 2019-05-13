package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	wg := sync.WaitGroup{}
	for i := 0; i <= 10000; i++ {
		wg.Add(1)
		go calc(&wg, i)
	}
	wg.Wait()
	fmt.Println("done")
}

func calc(w *sync.WaitGroup, i int) {
	fmt.Println("calc:", i)
	time.Sleep(time.Second * 2)
	w.Done()
}
