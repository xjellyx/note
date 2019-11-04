package main

import (
	"fmt"
	"github.com/golang/protobuf/proto"
	uuid "github.com/satori/go.uuid"
	"github.com/srlemon/note"
	"io/ioutil"
	"os"
)

func writer() {
	u1 := &note.Users{
		Uid:  *proto.String(uuid.NewV4().String()),
		Name: *proto.String("张三"),
	}
	//u2 := &note.Users{
	//	Uid:  uuid.NewV4().String(),
	//	Name: "李四",
	//}
	data, err := proto.Marshal(u1)
	if err != nil {
		panic(err)
	}
	ioutil.WriteFile("./test.txt", data, os.ModePerm)
}

func read() {
	//读取文件数据
	data, _ := ioutil.ReadFile("./test.txt")
	book := &note.Users{}
	//解码数据
	proto.Unmarshal(data, book)
	fmt.Println(book)
}

func main() {
}
