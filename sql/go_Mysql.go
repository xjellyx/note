package main

import (
	"encoding/json"
	"fmt"
	"github.com/Shopify/sarama"
	_ "github.com/go-sql-driver/mysql"
	"github.com/olongfen/note/log"
	"github.com/olongfen/note/sql/obj"
	uuid "github.com/satori/go.uuid"
	"reflect"
	"strings"
	"sync"
)

type student struct {
	Id   string `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
	Sex  string `json:"sex"`
}

var (
	wg = sync.WaitGroup{}
	l  = &sync.RWMutex{}
)

func main() {

	typeS := reflect.TypeOf(student{})
	switch typeS.Kind() {
	case reflect.Struct:
		num := typeS.NumField()
		for i := 0; i < num; i++ {
			fmt.Println(strings.ToLower(typeS.Field(i).Name))
		}
	}
	item := []string{"name", "age", "id", "sex"}

	d := strings.Join(item, ",")
	fmt.Println(d)
	var db *obj.DB
	var (
		s = new(student)
	)
	db = new(obj.DB)
	db.TableName = "student"
	db.Open()
	defer db.Close()

	str, err := db.ParamSQL("sql/test.sql")
	if err != nil {
		log.Println(err)
	}
	fmt.Println(str)

	// 建立表格
	if _, err = db.Create(str); err != nil {
		panic(err)
	}

	// 获取表的字段
	data, _err := db.Columns()
	if _err != nil {
		panic(err)
	}

	//  打印字段
	for _, v := range data {
		fmt.Println(v)
	}

	s.Name = "jack"
	s.Age = 18
	s.Id = uuid.NewV4().String()
	s.Sex = "boy"

	// 插入数据
	insert := fmt.Sprintf(`INSERT INTO %s(name,age,id,sex) VALUES ("%s",%d,"%s","%s")`, "student", s.Name, s.Age, s.Id, s.Sex)
	if err = db.Insert(insert); err != nil {
		panic(err)
	}

	// 更新数据
	update := fmt.Sprintf(`UPDATE %s SET name="%s" WHERE id="%s"`, "student", "Tom", s.Id)
	if err = db.Update(update); err != nil {
		panic(err)
	}
	// 查看数据
	query := fmt.Sprintf(`SELECT * FROM %s WHERE id="%s"`, "student", s.Id)
	rows, _err := db.QueryFind(query)
	if _err != nil {
		panic(_err)
	}
	for _, v := range rows {
		fmt.Println(v)
	}
	consumer, err := sarama.NewConsumer([]string{"localhost:9092"}, nil)

	if err != nil {
		panic(err)
	}

	partitionList, _err := consumer.Partitions("topic001")

	if _err != nil {
		panic(_err)
	}

	for partition := range partitionList {
		pc, err := consumer.ConsumePartition("topic001", int32(partition), sarama.OffsetNewest)
		if err != nil {
			panic(err)
		}

		defer pc.AsyncClose()

		wg.Add(1)

		go func(sarama.PartitionConsumer) {
			defer wg.Done()
			for msg := range pc.Messages() {
				fmt.Printf("Partition:%d, Offset:%d, Key:%s, Value:%s\n", msg.Partition, msg.Offset, string(msg.Key), string(msg.Value))
				if err = json.Unmarshal(msg.Value, s); err != nil {
					log.Errorln("aaaaaaaaaaaaaaa", err)
					continue
				}
				log.Infoln("sssssssssssssssss", s)
				insert = fmt.Sprintf(`INSERT INTO %s(name,age,id,sex) VALUES ("%s",%d,"%s","%s")`, "student", s.Name, s.Age, s.Id, s.Sex)
				if err = db.Insert(insert); err != nil {
					log.Errorln(err)
				}
			}
		}(pc)
		wg.Wait()
		consumer.Close()
	}
}
