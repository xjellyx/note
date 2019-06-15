package main

import (
	"context"
	"errors"
	"fmt"
	"github.com/apache/thrift/lib/go/thrift"
	"github.com/srlemon/note/thrift-rpc/gen-go/example"
	"reflect"
)

// 服务
type Sever struct {
}

const (
	HOST = "0.0.0.0" //
	PORT = "8898"
)

// 初始化数据,给后面需要时提供提供数据
var students []*example.Student

// 初始化数据
func init() {
	students = append(students, &example.Student{
		UID:       "fb189006-9a8b-4b50-a343-f221be1cce7b",
		Name:      "小明",
		Age:       18,
		ClassName: "一班",
		Sex:       "boy",
	})
	students = append(students, &example.Student{
		UID:       "a98ed979-881b-49a1-b52d-5747eedd3fe8",
		Name:      "小红",
		Age:       18,
		ClassName: "二班",
		Sex:       "girl",
	})
}

func (s *Sever) GetStudentByUID(ctx context.Context, uid string) (ret *example.Student, err error) {
	if len(uid) == 0 {
		err = errors.New("uid is nil")
		return
	}

	ret = new(example.Student)
	// 获取数据
	// (方法一)
	for _, v := range students {
		obj := reflect.ValueOf(v)
		for i := 0; i < obj.Elem().NumField(); i++ {
			if obj.Type().Field(i).Name == "UID" { // 获取UID属性
				if obj.Elem().Field(i).String() == uid { // 判断uid是否存在
					ret = new(example.Student)
					ret = v
					break
				} else {
					continue // uid不存在,跳出循环
				}
			} else {
				continue // 不是uid属性跳出循环
			}
		}
	}

	// 其实上面可以有另一种写法,我只是想摸索reflect包才那样写的
	// (方法二) 很简便的,比上面的逻辑简单很多,你可以注释上面的代码试一试这个
	//for _, v := range students {
	//	if v.UID == uid {
	//		ret = new(example.Student)
	//		ret = v
	//	} else {
	//		continue
	//	}
	//}

	if ret == nil {
		err = errors.New("student is not exist")
	}
	return
}

// 修改学生信息
func (s *Sever) ModifyStudent(ctx context.Context, uid string, form *example.FormStudent) (ret *example.Student, err error) {
	if len(uid) == 0 {
		err = errors.New("uid is nil")
		return
	} else if form == nil {
		err = errors.New("param is not valid")
	}

	var (
		data *example.Student
		num  int
	)

	// 判断是否存在该学生
	for _, v := range students {
		if v.UID == uid {
			data = new(example.Student)
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

	return
}
func main() {
	var (
		handler          = new(Sever) //
		processor        *example.BaseServiceProcessor
		serveTransport   *thrift.TServerSocket
		transportFactory thrift.TTransportFactory
		protocolFactory  *thrift.TBinaryProtocolFactory
		err              error
	)
	// new 一个处理器
	processor = example.NewBaseServiceProcessor(handler)
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
