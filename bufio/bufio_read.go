package main

import (
	"bufio"
	"fmt"
	"strings"
)

func main() {
	str := strings.NewReader("ABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890")
	fmt.Println(str)

	// 将 str 封装成一个带缓存的 bufio.Reader 对象
	buf := bufio.NewReaderSize(str, 0)
	var b = make([]byte, 10)

	// 返回缓存中未读取的数据的长度。
	fmt.Println(buf.Buffered())

	//  Peek 返回缓存的一个切片，该切片引用缓存中前 n 个字节的数据,该操作不会将数据读出，只是引用，
	//  引用的数据在下一次读取操作之 前是有效的。如果切片长度小于 n，则返回一个错误信息说明原因。
	//  如果 n 大于缓存的总大小，则返回 ErrBufferFull
	fmt.Println(buf.Peek(5))

	// Discard 跳过后续的 n 个字节的数据，返回跳过的字节数。
	// 如果结果小于 n，将返回错误信息。
	// 如果 n 小于缓存中的数据长度，则不会从底层提取数据。
	fmt.Println(buf.Discard(1))

	// 读取数据
	for n, err := 0, error(nil); err == nil; {
		n, err = buf.Read(b)
		fmt.Printf("%d   %q   %v\n", buf.Buffered(), b[:n], err)
	}
}
