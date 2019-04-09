package main

import (
	"fmt"
	"log"
	"os"
	"strings"
	"time"
)

func main() {
	var (
		err error
	)
	// 获取当前路径
	if _dir, _err := os.Getwd(); _err != nil {
		panic(_err)
	} else {
		fmt.Printf("pwd is %s", _dir)
	}

	// 获取系统环境变量的值
	path := os.Getenv("GOPATH")
	fmt.Printf("path is %v", path)

	// 获取主机名
	if _hostName, _err := os.Hostname(); _err == nil {

		fmt.Printf("HOST Name is %v", _hostName)
	} else {
		panic(_err)
	}

	// 获取用户ID
	fmt.Println(os.Getuid(), "ID")

	// 获取用户有效ID
	fmt.Println(os.Geteuid(), "ID")

	// 获取数组ID
	fmt.Println(os.Getgid(), "GID")

	// 获取有效数组ID
	fmt.Println(os.Getegid(), "GID")

	// 获取进程ID
	fmt.Println(os.Getpid(), "PID")

	// 获取父进程ID
	fmt.Println(os.Getppid(), "PID")

	// 设置环境变量的值
	if err = os.Setenv("TEST", "test"); err != nil {
		panic(err)
	}

	// 改变当前工作目录
	if err = os.Chdir("/home/allen/gocode/src/github.com/LnFen/note"); err != nil {
		panic(err)
	}
	fmt.Println(os.Getwd())

	// 创建文件
	if f1, err := os.Create("./1.txt"); err != nil {
		fmt.Println(err)
	} else {
		defer f1.Close()
	}

	// 修改文件权限
	if err = os.Chmod("./1.txt", 0777); err != nil {
		panic(err)
	}

	//修改文件所有者
	/*if err = os.Chown("./1.txt", 0, 0); err != nil {
		panic(err)
	}*/
	// 修改文件的访问时间和修改时间
	if err = os.Chtimes("./1.txt", time.Now().Add(time.Hour), time.Now().Add(time.Hour)); err != nil {
		panic(err)
	}

	// 获取所有环境变量
	fmt.Println(strings.Join(os.Environ(), "\r\n"))

	// 把字符串中带${var}或$var替换成指定指符串
	fmt.Println(os.Expand("${1} ${2} ${3}", func(k string) string {
		mapp := map[string]string{
			"1": "111",
			"2": "222",
			"3": "333",
		}
		return mapp[k]
	}))
	// 删除
	if err = os.Remove("./1.txt"); err != nil {
		panic(err)
	}

	// 创建目录
	if err = os.Mkdir("abc", os.ModePerm); err != nil {
		panic(err)
	}

	// 创建多级目录
	if err = os.MkdirAll("abc/d/e/f", os.ModePerm); err != nil {
		panic(err)
	}

	// 删除文件或目录
	if err = os.Remove("abc/d/e/f"); err != nil {
		log.Panicln(err)
	}

	// 删除指定目录下所有文件
	if err = os.RemoveAll("abc"); err != nil {
		log.Panicln(err)
	}

	// 重命名文件
	if _, err = os.Create("./2.txt"); err != nil {
		panic(err)
	}
	if err = os.Rename("./2.txt", "./2_new.txt"); err != nil {
		log.Panicln(err)
	}

	// 删除
	if err = os.Remove("./2_new.txt"); err != nil {
		panic(err)
	}
}
