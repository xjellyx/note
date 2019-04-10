package main

import "fmt"

// 这两个变量只能在该.go文件下访问
var a, b = 10, 20

// 这两个其他.go文件都可以访问
var A, B = 10, 20

// 全局变量范围>局部变量范围
// 代码块里面的变量只能在代码块里面访问使用，
func main() {
	var Num1 = 10
	{
		var num2 = 20
		// 代码块里面可以访问到Num1
		fmt.Println(Num1, num2)
	}
	// 代码块外面可以访问到Num1，但是访问不到num2
	fmt.Println(Num1)

	fmt.Println(addNum(a, b))
	fmt.Println(multiply(A, B))
}

func addNum(a, b int) int {
	return a + b
}

func multiply(a, b int) int {
	return a * b
}
