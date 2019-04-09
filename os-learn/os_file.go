package main

import (
	"fmt"
	"os"
	"strconv"
)

// 文件的创建与操作
func main() {
	var (
		file *os.File
		err  error
	)
	// 创建文件
	if file, err = os.Create("./test.txt"); err != nil {
		fmt.Println(err)
	}

	// 打开文件，不存在就创建
	if file, err = os.OpenFile("./test.txt", os.O_CREATE|os.O_RDWR|os.O_TRUNC, 0666); err != nil {
		fmt.Println(err)
	}

	// 打开文件
	if file, err = os.Open("./test.txt"); err != nil {
		fmt.Println(err)
	}

	// 关闭
	defer func() {
		if _err := recover(); _err != nil {
			panic(_err)
		} else {
			defer file.Close()
		}
	}()

	//   //修改文件权限，类似os.chmod
	if err = file.Chmod(0777); err != nil {
		fmt.Println(err)
	}

	// 返回文件的句柄，通过NewFile创建文件需要文件句柄
	fmt.Println(file.Fd(), "uintptr")

	if err = os.Remove("./test.txt"); err != nil {
		fmt.Println(err)
	}

	//向文件中写入数据
	if file, err = os.Create("./w.txt"); err != nil {
		fmt.Println(err)
	}
	for i := 0; i < 5; i++ {
		if _, err = file.Write([]byte("写入数据" + strconv.Itoa(i) + "\r\n")); err != nil {
			fmt.Println(err)
		}
	}

	os.Remove("./w.txt")
}
