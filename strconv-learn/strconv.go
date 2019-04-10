package main

import (
	"fmt"
	"strconv"
)

/* strconv 包下面常用的函数与方法*/
func main() {
	str := "hello\tworld\t 哈哈哈\n"
	fmt.Print(str)
	// 返回字符串在go语法下的双引号字面值表示，控制字符和不可打印字符会进行转义(\t,\n等)
	fmt.Println(strconv.Quote(str))

	// 返回字符串在go语法下的双引号字面值表示，除了上面的和非ASCII字符会进行转义
	fmt.Println(strconv.QuoteToASCII(str))

	// 返回字符串表示的bool值。它接受1、0、t，f、T、F、true、false、True、False、TRUE、FALSE；否则返回错误
	fmt.Println(strconv.ParseBool("T"))

	// 返回字符串表示的整数值，接受正负号。base指定进制(2到36),如果base为0，则会从字符串前置判断，”0x”代表16进制，
	// ”0”是8进制，否则是10进制；bitSize指定结果必须能无溢出的整数类型，0、8、16、32、64分别代表int，int8，int16，
	// int32，int64；返回的err是NumErr类型的，如果语法类型有误，err.Error=ErrSyntax，如果结果超出类型范围，
	// err.Error=ErrorRange
	fmt.Println(strconv.ParseInt("-100", 10, 0))

	// 类似ParseInt但不接受正负号，用于无符号整型
	fmt.Println(strconv.ParseUint("666", 0, 0))

	// 解析一个表示浮点数的字符串并返回其值。如果s合乎语法规则，函数会返回最为接近s表示值的一个浮点数(使用IEEE754规范舍入)。
	// bitSize指定了期望的接收类型，32是float32,64是float64,返回值是NumErr
	fmt.Println(strconv.ParseFloat("-3.14152874", 32))

	// 根据b的值返回”true”或”false”
	fmt.Println(strconv.FormatBool(true))

	// 返回的i的base进制的字符串表示，base必须在2-36之间，结果中会使用小写字母a到z表示大于10的数字
	// 整型转字符
	fmt.Println(strconv.FormatInt(100, 10)) // 把100 转为“100”

	// strconv.FormatUint(i uint64, base int)string -> 是FormatInt的无符号整数版本

	//  函数将浮点数表示为字符串并返回
	//fmt.Println(strconv.FormatFloat(3.141564, 10, 10, 64))

	// string 转 int
	fmt.Println(strconv.Atoi("100"))
	// int转string
	fmt.Println(strconv.Itoa(165))
}
