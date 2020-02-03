package main

import (
	"context"
	"errors"
	"fmt"
	"github.com/apache/thrift/lib/go/thrift"
	"github.com/olefen/note/thrift-rpc/gen-go/demo"
)

// 服务
type Sever struct {
}

const (
	HOST = "0.0.0.0" //
	PORT = "8898"
)

func main() {
	var (
		handler          = new(Sever) //
		processor        *demo.BaseServiceProcessor
		serveTransport   *thrift.TServerSocket
		transportFactory thrift.TTransportFactory
		protocolFactory  *thrift.TBinaryProtocolFactory
		err              error
	)
	// new 一个处理器
	processor = demo.NewBaseServiceProcessor(handler)
	// 创建服务
	if serveTransport, err = thrift.NewTServerSocket(HOST + ":" + PORT); err != nil {
		panic(err)
	}
	// 开机工厂模式
	transportFactory = thrift.NewTFramedTransportFactory(thrift.NewTTransportFactory())
	protocolFactory = thrift.NewTBinaryProtocolFactoryDefault()

	// 开启服务
	serve := thrift.NewTSimpleServer4(processor, serveTransport, transportFactory, protocolFactory)
	fmt.Printf("Running at:%s", HOST+":"+PORT)
	if err = serve.Serve(); err != nil {
		panic(err)
	}

}

func (s *Sever) GetStudentByUID(ctx context.Context, uid string) (ret *demo.Student, err error) {
	if len(uid) == 0 {
		err = errors.New("uid is nil")
		return
	}

	// 获取数据
	for _, v := range students {
		if v.UID == uid {
			ret = new(demo.Student)
			ret = v
		} else {
			continue
		}
	}

	// 判断是否有数据
	if ret == nil {
		err = errors.New("student is not exist")
		return
	} else if ret != nil && len(ret.UID) == 0 {
		err = errors.New("student is not exist")
	}
	return
}

// 修改学生信息
func (s *Sever) ModifyStudent(ctx context.Context, uid string, form *demo.FormStudent) (ret *demo.Student, err error) {
	if len(uid) == 0 {
		err = errors.New("uid is nil")
		return
	} else if form == nil {
		err = errors.New("param is not valid")
	}

	var (
		data *demo.Student
		num  int
	)

	// 判断是否存在该学生
	for _, v := range students {
		if v.UID == uid {
			data = new(demo.Student)
			data = v
		} else {
			continue
		}
	}
	// 不存在,返回
	if data == nil {
		err = errors.New("student is not exist")
	}
	if len(form.Name) > 0 {
		data.Name = form.Name
		num++
	}
	if len(form.ClassName) > 0 {
		data.ClassName = form.ClassName
		num++
	}
	if len(form.Sex) > 0 {
		data.Sex = form.Sex
		num++
	}
	if form.Age > 0 {
		data.Age = form.Age
		num++
	}

	// 更新数据,只能修改自己的数据
	for i := range students {
		if students[i].UID == uid {
			students[i] = data
			ret = students[i]
		}
	}

	// 判断是否有数据
	if ret == nil {
		err = errors.New("student is not exist")
		return
	} else if ret != nil && len(ret.UID) == 0 {
		err = errors.New("student is not exist")
	}

	return
}

func (s *Sever) GetData(ctx context.Context, data *demo.Data) (ret *demo.Data, err error) {

	ret = data

	return
}

var students []*demo.Student

func init() {
	students = append(students, &demo.Student{
		UID:       "fb189006-9a8b-4b50-a343-f221be1cce7b",
		Name:      "小明",
		Age:       18,
		ClassName: "一班",
		Sex:       "boy",
	})
	students = append(students, &demo.Student{
		UID:       "a98ed979-881b-49a1-b52d-5747eedd3fe8",
		Name:      "小红",
		Age:       18,
		ClassName: "二班",
		Sex:       "girl",
	})
}
