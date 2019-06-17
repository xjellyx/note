package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/coreos/etcd/clientv3"
	"time"
)

type user struct {
	Name string
	Age  int
}

func main() {
	cli, err := clientv3.New(
		clientv3.Config{
			Endpoints:   []string{"localhost:2379", "localhost:32774"},
			DialTimeout: 5 * time.Second,
		})
	if err != nil {
		panic(err)
	}
	defer cli.Close()
	go func() {
		rch := cli.Watch(context.Background(), "", clientv3.WithPrefix())
		for v := range rch {
			for _, ev := range v.Events {
				fmt.Println(string(ev.Kv.Key), string(ev.Kv.Value))
			}
		}
	}()
	time.Sleep(time.Second)
	// 设置 key1 的值为 value1
	a := user{
		Name: "Tom",
		Age:  18,
	}
	k1 := "key1"
	d, _ := json.Marshal(a)
	if resp, err := cli.Put(context.TODO(), k1, string(d), clientv3.WithPrevKV()); err != nil {
		println(err)
	} else {
		fmt.Println(resp)
	}

	//  // 设置 key1 的值为 value2, 并返回前一个值
	v2 := "value2"
	if resp, err := cli.Put(context.TODO(), k1, v2, clientv3.WithPrevKV()); err != nil {
		panic(err)
	} else {
		fmt.Println(resp)
	}

	v3 := "value3"
	if resp, err := cli.Put(context.TODO(), k1, v3, clientv3.WithPrevKV()); err != nil {
		panic(err)
	} else {
		fmt.Println(resp)
	}

}
