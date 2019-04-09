package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"strings"
)

func main() {
	var (
		data   []byte
		reader *bufio.Reader
		writer *bufio.Writer
		buf    io.Writer
		err    error
	)
	// 构造一个reader
	inputReader := strings.NewReader("1234567890")
	reader = bufio.NewReader(inputReader)

	// 构造一个writer
	buf = bytes.NewBuffer(make([]byte, 0))
	writer = bufio.NewWriter(buf)

	// 函数Peek函数: 返回缓存的一个Slice(引用,不是拷贝)，引用缓存中前n字节数据
	// > 如果引用的数据长度小于 n，则返回一个错误信息
	// > 如果 n 大于缓存的总大小，则返回 ErrBufferFull
	// 通过Peek的返回值，可以修改缓存中的数据, 但不能修改底层io.Reader中的数据

	if data, err = reader.Peek(5); err != nil {
		panic(err)
	}

	// 修改第一个字符
	data[0] = 'A'

	// 重新读取
	data, _ = reader.Peek(5)
	if _, err = writer.Write(data); err != nil {
		panic(err)
	}
	if err = writer.Flush(); err != nil {
		panic(err)
	}
	fmt.Println("buf(Changed): ", buf, "\ninputReadBuf(Not Changed): ", inputReader)

}
