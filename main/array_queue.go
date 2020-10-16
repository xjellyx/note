package main

import (
	"errors"
	"fmt"
)

var (
	ErrNilQueue = errors.New("arrQueue is nil")
)

type arrQueue []int

// Enqueue 队列尾部插入数据
func (q *arrQueue) Enqueue(e int) {
	*q = append(*q, e)
}

// Dequeue 队列删除第一个元素
func (q *arrQueue) Dequeue() (ret int, err error) {
	if len(*q) == 0 {
		err = ErrNilQueue
		return
	}
	temp := *q
	ret = temp[0]
	temp = temp[1:]
	*q = temp
	return
}

// First 返回第一个数据
func (q *arrQueue) First() (ret int, err error) {
	if len(*q) == 0 {
		err = ErrNilQueue
		return
	}
	temp := *q
	ret = temp[0]
	return
}

func (q *arrQueue) IsEmpty() bool {
	return len(*q) == 0
}

func (q *arrQueue) Len() int {
	return len(*q)
}

func main() {
	q := new(arrQueue)
	q.Enqueue(1)
	q.Enqueue(10)
	q.Enqueue(20)
	fmt.Println(q.First())   // 1
	fmt.Println(q.Dequeue()) // 1
	fmt.Println(q.First())   // 10
	fmt.Println(q.IsEmpty()) // false
	fmt.Println(q.Len())     // 2
	fmt.Println(q.Dequeue()) //10
	fmt.Println(q.Len())     // 1
}
