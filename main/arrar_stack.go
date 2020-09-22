package main

import (
	"errors"
	"fmt"
)

var (
	stackIsNil = errors.New("empty stack")
)

type stack []int

// Push 往stack顶部插入数据
func (s *stack) Push(e int) {
	*s = append(*s, e)
	return
}

// Pop 删除stack顶部数据并且返回删除的数据
func (s *stack) Pop() (ret int, err error) {
	if len(*s) == 0 {
		return 0, stackIsNil
	}

	temp := *s
	ret = temp[len(temp)-1]
	temp = temp[:len(temp)-1]
	*s = temp
	return
}

// IsEmpty 判断是否为空
func (s *stack) IsEmpty() bool {
	return len(*s) == 0
}

// Top 获取stack顶部数据
func (s *stack) Top() (int, error) {
	if len(*s) == 0 {
		return 0, stackIsNil
	}
	temp := *s
	return temp[len(temp)-1], nil
}

// Len 获取stack长度
func (s *stack) Len() int {
	return len(*s)
}

func main() {
	s := new(stack)
	// 插入1
	s.Push(1)
	// 插入2
	s.Push(2)
	// 插入5
	s.Push(5)
	// 获取长度
	fmt.Println(s.Len()) // 3
	// 获取stack顶部数据
	fmt.Println(s.Top()) // 5
	// 删除顶部数据
	fmt.Println(s.Pop()) // 5
	// 获取长度
	fmt.Println(s.Len()) // 2
	// 判断是否为空stack
	fmt.Println(s.IsEmpty())
}
